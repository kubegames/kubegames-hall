package room

import (
	"context"
	"encoding/json"
	"net/http"

	playerModel "github.com/kubegames/kubegames-hall/app/model/player"
	roomModel "github.com/kubegames/kubegames-hall/app/model/room"
	service "github.com/kubegames/kubegames-hall/app/service/room"
	"github.com/kubegames/kubegames-hall/app/service/room/types"
	"github.com/kubegames/kubegames-hall/internal/pkg/log"
	"github.com/kubegames/kubegames-hall/internal/pkg/nsq"
	mysql "github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
	"github.com/kubegames/kubegames-hall/internal/pkg/redis"
	"github.com/kubegames/kubegames-hall/pkg/platform"
	"github.com/kubegames/kubegames-hall/pkg/token"
	"xorm.io/xorm"
)

const (
	Namespace            = "games"
	GameBroadcastMessage = "GAME_BROADCAST_MESSAGE"
)

type roomImpl struct {
	mysql *mysql.Mysql
	redis redis.Redis
	nsq   nsq.Nsq
}

//new room
func NewRoom(mysql *mysql.Mysql, nsq nsq.Nsq) *roomImpl {
	//sync table
	if err := mysql.GetWriteEngine().Sync2(
		new(roomModel.Room),
		new(roomModel.RoomPool),
	); err != nil {
		panic(err.Error())
	}

	impl := &roomImpl{
		mysql: mysql,
		nsq:   nsq,
	}

	go impl.nsq.Consumer(platform.PlayerRecordMessage, "room", impl.UpdateRoomPool)

	return impl
}

func (impl *roomImpl) UpdateRoomPool(msg []byte) error {
	var records []*playerModel.PlayerRecord

	//session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//json unmarshal
	json.Unmarshal(msg, &records)

	//update room pool
	for _, record := range records {
		room, ok, err := roomModel.FindRoom(session, func(s *xorm.Session) *xorm.Session {
			return s.Where("room_id = ?", record.RoomID)
		})
		if err != nil {
			log.Warnf("get rooms %d error %s", record.RoomID, err.Error())
			continue
		}
		if !ok {
			log.Warnf("get rooms %d not found", record.RoomID)
			continue
		}
		if room.RoomPoolID > 0 && record.ProfitAmount != 0 {
			if _, err := roomModel.UpdateRoomPool(session, map[string]interface{}{}, func(s *xorm.Session) *xorm.Session {
				return s.Incr("water_level_line", record.ProfitAmount).ID(room.RoomPoolID)
			}); err != nil {
				log.Warnf("update room pool %d error %s", room.RoomPoolID, err.Error())
				continue
			}
		}
	}
	return nil
}

func (impl *roomImpl) AddGameRoom(ctx context.Context, request *types.AddGameRoomRequest) (response *types.AddGameRoomResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//insert room
	if _, err := roomModel.InsertRoom(session, roomModel.Room{
		PlatformID:               playerInfo.PlatformID,
		Status:                   false,
		RoomID:                   request.RoomID,
		GameID:                   request.GameID,
		Level:                    request.Level,
		Name:                     request.Name,
		Rate:                     request.Rate,
		EntranceRestrictions:     request.EntranceRestrictions,
		BottomNote:               request.BottomNote,
		AdviceConfig:             request.AdviceConfig,
		PointStatus:              request.PointStatus,
		MaxPeople:                request.MaxPeople,
		MinPeople:                request.MinPeople,
		IsAllowClose:             request.IsAllowClose,
		IsOpenAiRobot:            request.IsOpenAiRobot,
		IsOpenCrossPlatformMatch: request.IsOpenCrossPlatformMatch,
		AllowPlatformID:          request.AllowPlatformID,
		IsAllowAutoCreateTable:   request.IsAllowAutoCreateTable,
		Robot:                    request.Robot,
		RobotMaxBalance:          int64(request.RobotMaxBalance),
		RobotMinBalance:          int64(request.RobotMinBalance),
	}); err != nil {
		log.Errorf("mysql insert error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}

	return &types.AddGameRoomResponse{Success: true}, nil
}

func (impl *roomImpl) AddGameRooms(ctx context.Context, request *types.AddGameRoomsRequest) (response *types.AddGameRoomsResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		log.Errorf("mysql begin error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}

	for _, room := range request.Rooms {
		//insert room
		if _, err := roomModel.InsertRoom(session, roomModel.Room{
			PlatformID:               playerInfo.PlatformID,
			Status:                   false,
			RoomID:                   room.RoomID,
			GameID:                   room.GameID,
			Level:                    room.Level,
			Name:                     room.Name,
			Rate:                     room.Rate,
			EntranceRestrictions:     room.EntranceRestrictions,
			BottomNote:               room.BottomNote,
			AdviceConfig:             room.AdviceConfig,
			PointStatus:              room.PointStatus,
			MaxPeople:                room.MaxPeople,
			MinPeople:                room.MinPeople,
			IsAllowClose:             room.IsAllowClose,
			IsOpenAiRobot:            room.IsOpenAiRobot,
			IsOpenCrossPlatformMatch: room.IsOpenCrossPlatformMatch,
			AllowPlatformID:          room.AllowPlatformID,
			IsAllowAutoCreateTable:   room.IsAllowAutoCreateTable,
			Robot:                    room.Robot,
			RobotMaxBalance:          int64(room.RobotMaxBalance),
			RobotMinBalance:          int64(room.RobotMinBalance),
		}); err != nil {
			log.Errorf("mysql insert error %s", err.Error())
			return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
		}
	}

	//commit
	if err := session.Commit(); err != nil {
		log.Errorf("mysql commit insert error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.AddGameRoomsResponse{Success: true}, nil
}

func (impl *roomImpl) OpenGameRoom(ctx context.Context, request *types.OpenGameRoomRequest) (response *types.OpenGameRoomResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//update room
	if _, err := roomModel.UpdateRoom(session, map[string]interface{}{
		"status": true,
	}, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_id = ? AND platform_id = ?", request.RoomID, playerInfo.PlatformID)
	}); err != nil {
		log.Errorf("update rooms error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.OpenGameRoomResponse{Success: true}, nil
}

func (impl *roomImpl) CloseGameRoom(ctx context.Context, request *types.CloseGameRoomRequest) (response *types.CloseGameRoomResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//update room
	if _, err := roomModel.UpdateRoom(session, map[string]interface{}{
		"status": false,
	}, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_id = ? AND platform_id = ?", request.RoomID, playerInfo.PlatformID)
	}); err != nil {
		log.Errorf("update rooms error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.CloseGameRoomResponse{Success: true}, nil
}

func (impl *roomImpl) FindGameRoom(ctx context.Context, request *types.FindGameRoomRequest) (response *types.FindGameRoomResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetReadEngine().NewSession()
	defer session.Close()

	//find
	room, ok, err := roomModel.FindRoom(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_id = ? AND platform_id = ?", request.RoomID, playerInfo.PlatformID)
	})
	if err != nil {
		log.Errorf("get rooms %d error %s", request.RoomID, err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		return nil, service.NewRoomServiceError(http.StatusNotFound, "not found room")
	}
	return &types.FindGameRoomResponse{Room: &room}, nil
}

func (impl *roomImpl) FindGameRooms(ctx context.Context, request *types.FindGameRoomsRequest) (response *types.FindGameRoomsResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetReadEngine().NewSession()
	defer session.Close()

	//finds
	rooms, err := roomModel.FindRooms(session, func(s *xorm.Session) *xorm.Session {
		if request.Status == 1 {
			s = s.Where("status = ?", false)
		}
		if request.Status == 2 {
			s = s.Where("status = ?", true)
		}
		return s.Where("game_id = ? AND platform_id = ?", request.GameID, playerInfo.PlatformID)
	})
	if err != nil {
		log.Errorf("find gams error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.FindGameRoomsResponse{Rooms: rooms}, nil
}

func (impl *roomImpl) AddGameRoomPools(ctx context.Context, request *types.AddGameRoomPoolsRequest) (response *types.AddGameRoomPoolsResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//insert roompool
	var roompools []*roomModel.RoomPool
	for _, name := range request.Names {
		roompools = append(roompools, &roomModel.RoomPool{
			Name:       name,
			PlatformID: playerInfo.PlatformID,
		})
	}
	if len(roompools) > 0 {
		if _, err := roomModel.InsertRoomPools(session, roompools); err != nil {
			log.Errorf("inserts room pool error %s", err.Error())
			return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
		}
	}
	return &types.AddGameRoomPoolsResponse{Success: true}, nil
}

func (impl *roomImpl) BindGameRoomPool(ctx context.Context, request *types.BindGameRoomPoolRequest) (response *types.BindGameRoomPoolResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//find room pool
	roompool, ok, err := roomModel.FindRoomPool(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("id = ?", request.Id)
	})
	if err != nil {
		log.Errorf("find room pool %d error %s", request.Id, err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		return nil, service.NewRoomServiceError(http.StatusBadRequest, "room pool id error")
	}

	//update room
	if _, err := roomModel.UpdateRoom(session, map[string]interface{}{
		"room_pool_id": roompool.Id,
	}, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_id = ? AND platform_id = ?", request.RoomID, playerInfo.PlatformID)
	}); err != nil {
		log.Errorf("bind room pool %d to roomid %d error %s", request.Id, request.RoomID, err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.BindGameRoomPoolResponse{Success: true}, nil
}

func (impl *roomImpl) UnboundGameRoomPool(ctx context.Context, request *types.UnboundGameRoomPoolRequest) (response *types.UnboundGameRoomPoolResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//update room
	if _, err := roomModel.UpdateRoom(session, map[string]interface{}{
		"room_pool_id": 0,
	}, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_id = ? AND platform_id = ? AND room_pool_id = ?", request.RoomID, playerInfo.PlatformID, request.Id)
	}); err != nil {
		log.Errorf("unbound room pool %d to roomid %d error %s", request.Id, request.RoomID, err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.UnboundGameRoomPoolResponse{Success: true}, nil
}

func (impl *roomImpl) DeleteGameRoomPool(ctx context.Context, request *types.DeleteGameRoomPoolRequest) (response *types.DeleteGameRoomPoolResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		log.Errorf("begin error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}

	_, ok, err := roomModel.FindRoomPool(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_pool_id = ? AND platform_id = ?", request.Id, playerInfo.PlatformID)
	})
	if err != nil {
		log.Errorf("find room pool %d error %s", request.Id, err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		log.Errorf("room pool %d not found", request.Id)
		return nil, service.NewRoomServiceError(http.StatusNotFound, "room pool not found")
	}

	//unbound
	if _, err := roomModel.UpdateRoom(session, map[string]interface{}{
		"room_pool_id": 0,
	}, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_pool_id = ?", request.Id)
	}); err != nil {
		log.Errorf("unbound all room pool %d error %s", request.Id, err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}

	//delete
	if _, err := roomModel.DeleteRoomPool(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_pool_id = ?", request.Id)
	}); err != nil {
		log.Errorf("delete room pool %d error %s", request.Id, err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}

	//commit
	if err := session.Commit(); err != nil {
		log.Errorf("begin error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.DeleteGameRoomPoolResponse{Success: true}, nil
}

func (impl *roomImpl) FindGameRoomPool(ctx context.Context, request *types.FindGameRoomPoolRequest) (response *types.FindGameRoomPoolResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetReadEngine().NewSession()
	defer session.Close()

	roompool, ok, err := roomModel.FindRoomPool(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("id = ? AND platform_id = ?", request.Id, playerInfo.PlatformID)
	})
	if err != nil {
		log.Errorf("find room pool %d error %s", request.Id, err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		log.Errorf("room pool %d not found", request.Id)
		return nil, service.NewRoomServiceError(http.StatusNotFound, "room pool not found")
	}

	return &types.FindGameRoomPoolResponse{RoomPool: &roompool}, nil
}

func (impl *roomImpl) FindGameRoomPools(ctx context.Context, request *types.FindGameRoomPoolsRequest) (response *types.FindGameRoomPoolsResponse, err error) {
	playerInfo, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusUnauthorized, "permissions error")
	}

	//new session
	session := impl.mysql.GetReadEngine().NewSession()
	defer session.Close()

	roompools, err := roomModel.FindRoomPools(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("platform_id = ?", playerInfo.PlatformID)
	})
	if err != nil {
		log.Errorf("find room pools error %s", err.Error())
		return nil, service.NewRoomServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.FindGameRoomPoolsResponse{RoomPools: roompools}, nil
}
