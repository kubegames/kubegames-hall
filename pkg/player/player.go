package player

import (
	"context"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/kubegames/kubegames-hall/app/message"
	playerModel "github.com/kubegames/kubegames-hall/app/model/player"
	roomModel "github.com/kubegames/kubegames-hall/app/model/room"
	"github.com/kubegames/kubegames-hall/app/service/match"
	matchTypes "github.com/kubegames/kubegames-hall/app/service/match/types"
	service "github.com/kubegames/kubegames-hall/app/service/player"
	"github.com/kubegames/kubegames-hall/app/service/player/types"
	"github.com/kubegames/kubegames-hall/internal/pkg/log"
	"github.com/kubegames/kubegames-hall/internal/pkg/nats"
	mysql "github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
	"github.com/kubegames/kubegames-hall/pkg/platform"
	"github.com/kubegames/kubegames-hall/pkg/token"
	natsgo "github.com/nats-io/nats.go"
	"google.golang.org/grpc"
	"xorm.io/xorm"
)

//set rand
func init() {
	rand.Seed(time.Now().UnixNano())
}

//change websocket
var upGrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type (
	//player impl
	playerImpl struct {
		mysql *mysql.Mysql
		nats  nats.Nats
		list  sync.Map
	}
)

//new player
func NewPlayer(mysql *mysql.Mysql, nats nats.Nats) *playerImpl {
	impl := &playerImpl{
		mysql: mysql,
		nats:  nats,
	}

	//sync table
	if err := impl.mysql.GetWriteEngine().Sync2(
		new(playerModel.Player),
		new(playerModel.Status),
	); err != nil {
		panic(err.Error())
	}

	//subscribe message
	if _, err := impl.nats.Engine().Subscribe(platform.GameBroadcastMessage, impl.PushMessage); err != nil {
		panic(err)
	}

	//return
	return impl
}

//push message
func (impl *playerImpl) PushMessage(m *natsgo.Msg) {
	impl.list.Range(func(key, value interface{}) bool {
		conn := value.(*websocket.Conn)
		if err := conn.WriteMessage(websocket.BinaryMessage, m.Data); err != nil {
			log.Warnf("push to player %v message error %s", key, err.Error())
		}
		return true
	})
}

//player websocket connect
func (impl *playerImpl) Websocket(c *gin.Context) {
	//get token string
	tokenString := c.Query("token")

	//vcalidate token formate
	if tokenString == "" {
		log.Errorf("token is nil")
		return
	}

	playerinfo, err := token.GetPlayerByToken(tokenString)
	if err != nil {
		log.Errorf("token is error %s", err.Error())
		return
	}

	//websocket connect
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Errorf("upgrade error %s", err.Error())
		return
	}
	defer conn.Close()

	log.Tracef("player %d platform online", playerinfo.PlayerID)

	impl.list.Store(playerinfo.PlayerID, conn)

	log.Tracef("read player %d platform mesage", playerinfo.PlayerID)
	for {
		//read
		_, _, err := conn.ReadMessage()
		if err != nil {
			log.Errorf("read message error %s", err.Error())
			break
		}

		//proto
		message, err := proto.Marshal(&message.FrameMsg{MainCmd: 0, SubCmd: 2, Buff: nil, Time: time.Now().Unix()})
		if err != nil {
			log.Warnln("proto marshal error %s", err.Error())
			break
		}

		//write
		if err := conn.WriteMessage(websocket.BinaryMessage, message); err != nil {
			log.Warnf("send message error %s", err.Error())
			break
		}
	}

	impl.list.Delete(playerinfo.PlayerID)
	log.Tracef("player %d platform offline", playerinfo.PlayerID)
}

//player login
func (impl *playerImpl) PlayerLogin(ctx context.Context, request *types.PlayerLoginRequest) (response *types.PlayerLoginResponse, err error) {
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//get player info
	playerInfo, ok, err := playerModel.FindPlayer(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("phone = ?", request.Phone)
	})
	if err != nil {
		log.Errorf("get player phone %s error %s", request.Phone, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}

	//not found
	if !ok {
		playerInfo = playerModel.Player{
			Phone:      request.Phone,
			Balance:    100000000,
			Nick:       "player",
			Account:    request.Phone,
			PlatformID: 1,
			Sign:       "sigin",
			CreatedAt:  time.Now(),
			UpdatedAt:  time.Now(),
			Version:    1,
		}

		//insert
		id, err := impl.mysql.GetWriteEngine().Insert(playerInfo)
		if err != nil {
			log.Errorf("insert player phone %s error %s", request.Phone, err.Error())
			return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
		}

		//set player id
		playerInfo.PlayerID = uint32(id)
	} else {
		if playerInfo.Password != request.Password {
			return nil, service.NewPlayerServiceError(http.StatusBadRequest, "password error")
		}
	}

	//token
	token, err := token.NewPlayerToken(&playerInfo)
	if err != nil {
		log.Errorf("jwt marshal error %s", err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, "create token fail")
	}

	//create response
	return &types.PlayerLoginResponse{
		Token: token,
	}, nil
}

//player match
func (impl *playerImpl) playerMatch(ctx context.Context, adderss string, gameid, roomid uint32, player *playerModel.Player, token string) (*matchTypes.MatchResponse, error) {
	conn, err := grpc.Dial(adderss, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	rsp, err := match.NewMatchServiceClient(conn).Match(ctx, &matchTypes.MatchRequest{
		GameID: gameid,
		RoomID: roomid,
		Player: player,
		Token:  token,
	})
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

//player match
func (impl *playerImpl) PlayerMatch(ctx context.Context, request *types.PlayerMatchRequest) (response *types.PlayerMatchResponse, err error) {
	player, token, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewPlayerServiceError(http.StatusUnauthorized, "permissions error")
	}

	//session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	//get player info
	playerInfo, ok, err := playerModel.FindPlayer(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", player.PlayerID)
	})
	if err != nil {
		log.Errorf("get player %d info error %s", player.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		log.Errorf("get player %d info error %s", playerInfo.PlayerID, "player not found")
		return nil, service.NewPlayerServiceError(http.StatusNotFound, "player not found")
	}

	log.Tracef("start match player %d balance %d chip %d", playerInfo.PlayerID, playerInfo.Balance, playerInfo.Chip)

	//default response
	response = &types.PlayerMatchResponse{Success: false, Note: "match fail"}

	//find room
	room, ok, err := roomModel.FindRoom(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_id = ?", request.RoomID)
	})
	if err != nil {
		log.Errorf("get room %d error %s", request.RoomID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}
	if ok == false || room.Status == false {
		response.Note = "room is unavailable"
		return response, nil
	}

	//get player status
	status, ok, err := playerModel.FindStatus(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", playerInfo.PlayerID)
	})
	if err != nil {
		log.Errorf("get player %d status error %s", playerInfo.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}
	if ok {
		if len(status.PodName) > 0 || len(status.Ip) > 0 {
			response.Status = &status
			response.Note = "match success"
			response.Success = true
			return response, nil
		}
		if _, err := playerModel.DeleteStatus(session, func(s *xorm.Session) *xorm.Session {
			return s.Where("player_id = ?", playerInfo.PlayerID)
		}); err != nil {
			log.Errorf("delete player %d status error %s", playerInfo.PlayerID, err.Error())
			return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
		}
	}

	//check time
	runRooms, err := roomModel.FindServerRunRooms(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("room_id = ?", room.RoomID)
	})
	if err != nil {
		response.Note = "service is unavailable"
		return response, nil
	}

	for _, server := range runRooms {
		result, err := impl.playerMatch(ctx, server.Ip, request.GameID, request.RoomID, &playerInfo, token)
		if err != nil {
			log.Warnf("player match room %d servers %s error %s", server.RoomID, server.Ip, err.Error())
			continue
		}
		if result.Success == true {
			status.PodName = result.PodName
			status.Ip = result.PodIp
			status.Token = result.Token
			response.Status = &status
			response.Note = "match success"
			response.Success = true
			break
		}
		if result.Stop == false {
			response.Note = result.Note
			break
		}
	}
	if response.Success {
		if _, err := playerModel.InsertStatus(session, playerModel.Status{
			PlayerID: playerInfo.PlayerID,
			GameID:   request.GameID,
			RoomID:   request.RoomID,
			PodName:  response.Status.PodName,
			Ip:       response.Status.Ip,
			Token:    response.Status.Token,
		}); err != nil {
			log.Errorf("insert palyer %d status error %s", playerInfo.PlayerID, err.Error())
			response.Note = "match fail"
			return response, nil
		}
	}
	return response, nil
}

//player info
func (impl *playerImpl) PlayerInfo(ctx context.Context, request *types.PlayerInfoRequest) (response *types.PlayerInfoResponse, err error) {
	player, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewPlayerServiceError(http.StatusUnauthorized, "permissions error")
	}

	//session
	session := impl.mysql.GetReadEngine().NewSession()
	defer session.Close()

	//get player info
	playerInfo, ok, err := playerModel.FindPlayer(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", player.PlayerID)
	})
	if err != nil {
		log.Errorf("get player %d info error %s", player.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		log.Errorf("get player %d info error %s", playerInfo.PlayerID, "player not found")
		return nil, service.NewPlayerServiceError(http.StatusNotFound, "player not found")
	}

	response = &types.PlayerInfoResponse{PlayerInfo: &playerInfo}

	//get player status
	status, ok, err := playerModel.FindStatus(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", playerInfo.PlayerID)
	})
	if err != nil {
		log.Errorf("get player %d status error %s", playerInfo.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}

	if ok {
		if len(status.PodName) > 0 || len(status.Ip) > 0 {
			response.Status = &status
		}
	}
	return response, nil
}

//player give
func (impl *playerImpl) PlayerGive(ctx context.Context, request *types.PlayerGiveRequest) (response *types.PlayerGiveResponse, err error) {
	player, _, err := token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewPlayerServiceError(http.StatusUnauthorized, "permissions error")
	}

	//check balance
	if request.Balance <= 0 {
		return nil, service.NewPlayerServiceError(http.StatusBadRequest, "balance cannot be negative")
	}

	//session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		log.Errorf("mysql begin error %s", err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}

	givePlayer, ok, err := playerModel.FindPlayer(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", request.PlayerID)
	})
	if err != nil {
		log.Errorf("get player %d info error %s", request.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		log.Errorf("get player %d info error %s", request.PlayerID, "player not found")
		return nil, service.NewPlayerServiceError(http.StatusOK, "player not found")
	}

	playerInfo, ok, err := playerModel.FindPlayer(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", player.PlayerID)
	})
	if err != nil {
		log.Errorf("get player %d info error %s", player.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		log.Errorf("get player %d info error %s", player.PlayerID, "player not found")
		return nil, service.NewPlayerServiceError(http.StatusNotFound, "player not found")
	}
	if playerInfo.Balance-request.Balance < 0 {
		return nil, service.NewPlayerServiceError(http.StatusBadRequest, "sorry, your credit is running low")
	}

	//update player
	if _, err := playerModel.UpdatePlayer(session, map[string]interface{}{
		"balance": playerInfo.Balance - request.Balance,
	}, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", playerInfo.PlayerID).Where("balance = ?", playerInfo.Balance)
	}); err != nil {
		log.Errorf("update player %d error %s", playerInfo.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}

	//update give player
	if _, err := playerModel.UpdatePlayer(session, map[string]interface{}{
		"balance": givePlayer.Balance + request.Balance,
	}, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", givePlayer.PlayerID).Where("balance = ?", givePlayer.Balance)
	}); err != nil {
		log.Errorf("update give player %d error %s", givePlayer.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}

	//insert player bill
	if _, err := playerModel.InsertPlayerBill(session, playerModel.PlayerBill{
		PlayerID: playerInfo.PlayerID,
		Before:   playerInfo.Balance,
		Balance:  -request.Balance,
		After:    playerInfo.Balance - request.Balance,
		Kind:     playerModel.PlayerBill_Give,
	}); err != nil {
		log.Errorf("insert player %d bill error %s", playerInfo.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}

	//insert give player bill
	if _, err := playerModel.InsertPlayerBill(session, playerModel.PlayerBill{
		PlayerID: givePlayer.PlayerID,
		Before:   givePlayer.Balance,
		Balance:  request.Balance,
		After:    givePlayer.Balance + request.Balance,
		Kind:     playerModel.PlayerBill_Give,
	}); err != nil {
		log.Errorf("insert give player %d bill error %s", givePlayer.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}

	if err := session.Commit(); err != nil {
		log.Errorf("mysql commit error %s", err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.PlayerGiveResponse{Success: true}, nil
}

//player recharge
func (impl *playerImpl) PlayerRecharge(ctx context.Context, request *types.PlayerRechargeRequest) (response *types.PlayerRechargeResponse, err error) {
	_, _, err = token.GetPlayerByContext(ctx)
	if err != nil {
		log.Errorf("authentication error %s", err.Error())
		return nil, service.NewPlayerServiceError(http.StatusUnauthorized, "permissions error")
	}

	//session
	session := impl.mysql.GetWriteEngine().NewSession()
	defer session.Close()

	if err := session.Begin(); err != nil {
		log.Errorf("mysql begin error %s", err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}

	rechargePlayer, ok, err := playerModel.FindPlayer(session, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", request.PlayerID)
	})
	if err != nil {
		log.Errorf("get player %d info error %s", request.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}
	if !ok {
		log.Errorf("get player %d info error %s", request.PlayerID, "player not found")
		return nil, service.NewPlayerServiceError(http.StatusOK, "player not found")
	}

	//update player
	if _, err := playerModel.UpdatePlayer(session, map[string]interface{}{
		"balance": rechargePlayer.Balance + request.Balance,
	}, func(s *xorm.Session) *xorm.Session {
		return s.Where("player_id = ?", rechargePlayer.PlayerID).Where("balance = ?", rechargePlayer.Balance)
	}); err != nil {
		log.Errorf("update player %d error %s", rechargePlayer.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}

	//insert player bill
	if _, err := playerModel.InsertPlayerBill(session, playerModel.PlayerBill{
		PlayerID: rechargePlayer.PlayerID,
		Before:   rechargePlayer.Balance,
		Balance:  request.Balance,
		After:    rechargePlayer.Balance + request.Balance,
		Kind:     playerModel.PlayerBill_Recharge,
	}); err != nil {
		log.Errorf("insert player %d bill error %s", rechargePlayer.PlayerID, err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}

	if err := session.Commit(); err != nil {
		log.Errorf("mysql commit error %s", err.Error())
		return nil, service.NewPlayerServiceError(http.StatusInternalServerError, err.Error())
	}
	return &types.PlayerRechargeResponse{Success: true}, nil
}
