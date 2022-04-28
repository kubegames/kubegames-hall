package player

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kubegames/kubegames-hall/app/service/player/types"
	"github.com/kubegames/kubegames-hall/internal/pkg/nats"
	"github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
)

type PlayerService interface {

	//player websocket connect
	Websocket(c *gin.Context)

	//get phone code
	GetPhoneCode(ctx context.Context, request *types.GetPhoneCodeRequest) (response *types.GetPhoneCodeResponse, err error)

	//player register
	PlayerRegister(ctx context.Context, request *types.PlayerRegisterRequest) (response *types.PlayerRegisterResponse, err error)

	//player login
	PlayerLogin(ctx context.Context, request *types.PlayerLoginRequest) (response *types.PlayerLoginResponse, err error)

	//player match
	PlayerMatch(ctx context.Context, request *types.PlayerMatchRequest) (response *types.PlayerMatchResponse, err error)

	//player info
	PlayerInfo(ctx context.Context, request *types.PlayerInfoRequest) (response *types.PlayerInfoResponse, err error)

	//player give
	PlayerGive(ctx context.Context, request *types.PlayerGiveRequest) (response *types.PlayerGiveResponse, err error)

	//player recharge
	PlayerRecharge(ctx context.Context, request *types.PlayerRechargeRequest) (response *types.PlayerRechargeResponse, err error)
}

type Service struct {
	*playerImpl
}

func NewService(mysql *xorm.Mysql, nats nats.Nats) PlayerService {
	return &Service{
		playerImpl: NewPlayer(mysql, nats),
	}
}
