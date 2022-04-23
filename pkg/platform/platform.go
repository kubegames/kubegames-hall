package platform

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	playerModel "github.com/kubegames/kubegames-hall/app/model/player"
	roomModel "github.com/kubegames/kubegames-hall/app/model/room"
	"github.com/kubegames/kubegames-hall/app/service/match"
	matchTypes "github.com/kubegames/kubegames-hall/app/service/match/types"
	service "github.com/kubegames/kubegames-hall/app/service/platform"
	"github.com/kubegames/kubegames-hall/app/service/platform/types"
	"github.com/kubegames/kubegames-hall/internal/pkg/log"
	"github.com/kubegames/kubegames-hall/internal/pkg/nats"
	"github.com/kubegames/kubegames-hall/internal/pkg/nsq"
	mysql "github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
	"github.com/kubegames/kubegames-hall/pkg/token"
	"google.golang.org/grpc"
	"xorm.io/xorm"
)

const (
	GameBroadcastMessage = "GAME_BROADCAST_MESSAGE"
	PlayerRecordMessage  = "PLAYER_RECORD_MESSAGE"
)

//platform impl
type platformImpl struct {
	mysql *mysql.Mysql
	nats  nats.Nats
	nsq   nsq.Nsq
}

func NewPlatform(mysql *mysql.Mysql, nats nats.Nats, nsq nsq.Nsq) *platformImpl {
	//sync table
	if err := mysql.GetWriteEngine().Sync2(
		new(roomModel.ServerRunRoom),
		new(roomModel.ServerLockRoom),
		new(playerModel.PlayerRecord),
		new(playerModel.PlayerBill),
	); err != nil {
		panic(err.Error())
	}

	impl := &platformImpl{mysql: mysql, nats: nats, nsq: nsq}
	go impl.CheckServerRunRooms()
	return impl
}

//player match
func (impl *platformImpl) ping(ctx context.Context, address string) (bool, error) {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Errorf("did not connect: %v", err)
		return false, err
	}
	defer conn.Close()
	rsp, err := match.NewMatchServiceClient(conn).Ping(ctx, &matchTypes.PingRequest{})
	if err != nil {
		log.Errorf("grpc ping call error %s", err.Error())
		return false, err
	}
	return rsp.Success, nil
}

//检测房间过期释放锁
func (impl *platformImpl) CheckServerRunRooms() {
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()
	for {
		//sleep
		time.Sleep(5 * time.Second)

		//check time
		rooms, err := roomModel.FindServerRunRooms(session)
		if err != nil {
			log.Warnf("find lock expiration rooms error %s", err.Error())
			continue
		}

		for _, room := range rooms {
			if _, err := impl.ping(context.Background(), room.Ip); err != nil {
				log.Warnf("ping game %d server error %s", room.RoomID, err.Error())

				//释放玩家
				if _, err := playerModel.DeleteStatus(session, func(s *xorm.Session) *xorm.Session {
					return s.Where("ip = ?", room.Ip)
				}); err != nil {
					log.Errorf("delete ip %s player error %s", room.Ip, err.Error())
					continue
				}

				//删除运行的房间
				if _, err := roomModel.DeleteServerRunRoom(session, func(s *xorm.Session) *xorm.Session {
					return s.Where("room_id = ? AND ip = ?", room.RoomID, room.Ip)
				}); err != nil {
					log.Errorf("delete server run room error %s", err.Error())
					continue
				}

				//删除锁住的房间
				if _, err := roomModel.DeleteServerLockRoom(session, func(s *xorm.Session) *xorm.Session {
					return s.Where("room_id = ? AND ip = ?", room.RoomID, room.Ip)
				}); err != nil {
					log.Errorf("delete server lock room error %s", err.Error())
					continue
				}
			}
		}
	}
}

//广播
func (impl *platformImpl) Broadcast(ctx context.Context, request *types.BroadcastRequest) (response *types.BroadcastResponse, err error) {
	if err := impl.nats.Engine().Publish(GameBroadcastMessage, request.Buff); err != nil {
		log.Errorf("broadcast publish message error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.BroadcastResponse{Success: true}, nil
}

//玩家离开游戏
func (impl *platformImpl) PlayerLeaveGame(ctx context.Context, request *types.PlayerLeaveGameRequest) (response *types.PlayerLeaveGameResponse, err error) {
	if request.IsRobot {
		log.Tracef("robot %d leave game", request.PlayerID)
		return &types.PlayerLeaveGameResponse{Success: true}, nil
	}

	//session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//delete
	if _, err := playerModel.DeleteStatus(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", request.PlayerID)
	}); err != nil {
		log.Errorf("delete playerid %d status error %s", request.PlayerID, err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}
	log.Tracef("player %d leave game", request.PlayerID)
	return &types.PlayerLeaveGameResponse{Success: true}, nil
}

//申请机器人
func (impl *platformImpl) ApplyRobot(ctx context.Context, request *types.ApplyRobotRequest) (response *types.ApplyRobotResponse, err error) {
	var index uint32 = 0

	rsp := &types.ApplyRobotResponse{
		GameID: request.GameID,
		RoomID: request.RoomID,
	}

	for index = 0; index < request.Number; index++ {
		id := uint32(time.Now().Nanosecond())

		//mock platform
		playerInfo := &playerModel.Player{
			PlayerID:   id,
			IsRobot:    true,
			Phone:      fmt.Sprintf("%d", id),
			Balance:    rand.Int63n(request.MaxBalance-request.MinBalance) + request.MinBalance,
			Nick:       "Robot",
			PlatformID: 1,
		}

		//token
		token, err := token.NewPlayerToken(playerInfo)
		if err != nil {
			log.Errorf("jwt marshal error %s", err.Error())
			return nil, service.NewPlatformServiceError(http.StatusInternalServerError, "create token fail")
		}

		robotInfo := &types.RobotInfo{
			Player: playerInfo,
			Token:  token,
		}
		rsp.List = append(rsp.List, robotInfo)
	}

	return rsp, nil
}

//获取房间列表
func (impl *platformImpl) FindRooms(ctx context.Context, request *types.FindRoomsRequest) (response *types.FindRoomsResponse, err error) {
	//session
	session := impl.mysql.GetReadEngine().NewSession()
	defer session.Close()

	//find rooms
	rooms, err := roomModel.FindRooms(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("game_id = ? AND status = ?", request.GameID, true)
	})
	if err != nil {
		log.Errorf("find gams error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.FindRoomsResponse{Rooms: rooms}, nil
}

//争取房间
func (impl *platformImpl) AcquireRoom(ctx context.Context, request *types.AcquireRoomRequest) (response *types.AcquireRoomResponse, err error) {
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//开启数据库事务
	if err := session.Begin(); err != nil {
		log.Errorf("mysql begin error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}

	//插入数据成功表示抢占成功
	if _, err := roomModel.InsertServerLockRoom(session, roomModel.ServerLockRoom{RoomID: request.RoomID, Ip: request.Ip, CreatedAt: time.Now()}); err != nil {
		log.Errorf("insert server lock room error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}

	//开启房间
	if _, err := roomModel.InsertServerRunRoom(session, roomModel.ServerRunRoom{RoomID: request.RoomID, Ip: request.Ip}); err != nil {
		log.Errorf("insert server run room error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}

	//合并
	if err := session.Commit(); err != nil {
		log.Errorf("mysql commit insert error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.AcquireRoomResponse{Success: true}, nil
}

//归还房间
func (impl *platformImpl) ReleaseRoom(ctx context.Context, request *types.ReleaseRoomRequest) (response *types.ReleaseRoomResponse, err error) {
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//开启数据库事务
	if err := session.Begin(); err != nil {
		log.Errorf("mysql begin error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}

	//删除锁住的房间
	if _, err := roomModel.DeleteServerLockRoom(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_id = ? AND ip = ?", request.RoomID, request.Ip)
	}); err != nil {
		log.Errorf("delete server lock room error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}

	//删除运行的房间
	if _, err := roomModel.DeleteServerRunRoom(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_id = ? AND ip = ?", request.RoomID, request.Ip)
	}); err != nil {
		log.Errorf("delete server run room error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}

	//合并
	if err := session.Commit(); err != nil {
		log.Errorf("mysql commit insert error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}

	return &types.ReleaseRoomResponse{Success: true}, nil
}

//启动房间
func (impl *platformImpl) RunningRoom(ctx context.Context, request *types.RunningRoomRequest) (response *types.RunningRoomResponse, err error) {
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	_, ok, err := roomModel.FindServerRunRoom(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_id = ?", request.RoomID).Where("ip = ?", request.Ip)
	})
	if err != nil {
		log.Errorf("find server run room error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}

	//开启房间
	if !ok {
		if _, err := roomModel.InsertServerRunRoom(session, roomModel.ServerRunRoom{RoomID: request.RoomID, Ip: request.Ip}); err != nil {
			log.Errorf("insert server run room error %s", err.Error())
			return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
		}
	}
	return &types.RunningRoomResponse{Success: true}, nil
}

//upload player record
func (impl *platformImpl) UploadPlayerRecord(ctx context.Context, request *types.UploadPlayerRecordRequest) (response *types.UploadPlayerRecordResponse, err error) {
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()
	if message, err := json.Marshal(&request.Records); err != nil {
		log.Warnf("json marshal player record error %s", err.Error())
	} else {
		if err := impl.nsq.Publish(PlayerRecordMessage, message); err != nil {
			log.Warnf("nsq publish player record error %s", err.Error())
		}
	}

	if _, err := playerModel.InsertPlayerRecords(session, request.Records); err != nil {
		log.Errorf("upload player record error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.UploadPlayerRecordResponse{Success: true}, nil
}

//upload player settle info
func (impl *platformImpl) UploadPlayerSettleInfo(ctx context.Context, request *types.UploadPlayerSettleInfoRequest) (response *types.UploadPlayerSettleInfoResponse, err error) {
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		log.Errorf("mysql begin error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}

	playerInfo, ok, err := playerModel.FindPlayer(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", request.PlayerID)
	})
	if err != nil {
		log.Errorf("get player %d info error %s", request.PlayerID, err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		log.Errorf("get player %d info error %s", request.PlayerID, "player not found")
		return nil, service.NewPlatformServiceError(http.StatusOK, "player not found")
	}

	//update player
	if _, err := playerModel.UpdatePlayer(session, map[string]interface{}{
		"balance": playerInfo.Balance + request.Balance,
		"chip":    request.Chip,
	}, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", playerInfo.PlayerID).Where("balance = ?", playerInfo.Balance)
	}); err != nil {
		log.Errorf("update player %d error %s", playerInfo.PlayerID, err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}

	//insert player bill
	if _, err := playerModel.InsertPlayerBill(session, playerModel.PlayerBill{
		PlayerID: playerInfo.PlayerID,
		Before:   playerInfo.Balance,
		Balance:  request.Balance,
		After:    playerInfo.Balance + request.Balance,
		Kind:     playerModel.PlayerBill_Game,
	}); err != nil {
		log.Errorf("insert player %d bill error %s", playerInfo.PlayerID, err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}

	if err := session.Commit(); err != nil {
		log.Errorf("mysql commit error %s", err.Error())
		return nil, service.NewPlatformServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.UploadPlayerSettleInfoResponse{Success: true}, nil
}

func (impl *platformImpl) GetRoomPool(ctx context.Context, request *types.GetRoomPoolRequest) (response *types.GetRoomPoolResponse, err error) {
	session := impl.mysql.GetReadEngine().NewSession()
	defer session.Close()

	if room, ok, err := roomModel.FindRoom(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_id = ?", request.RoomID)
	}); err == nil && ok {
		if room.RoomPoolID > 0 {
			if roompool, ok, err := roomModel.FindRoomPool(session, func(s *xorm.Session) *xorm.Session {
				return s.Where("room_id = ?", request.RoomID)
			}); err == nil && ok {
				return &types.GetRoomPoolResponse{Value: roompool.WaterLevelLine}, nil
			}
		}
	}
	return &types.GetRoomPoolResponse{Value: 0}, nil
}
