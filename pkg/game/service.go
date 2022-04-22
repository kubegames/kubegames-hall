package game

import (
	"context"

	"github.com/kubegames/kubegames-hall/app/service/game/types"
	"github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
	gamesclientset "github.com/kubegames/kubegames-operator/pkg/client/game/clientset/versioned"
	"k8s.io/client-go/kubernetes"
)

type GameService interface {
	AddGameRequest(ctx context.Context, request *types.AddGameRequest) (response *types.AddGameResponse, err error)

	StartGame(ctx context.Context, request *types.StartGameRequest) (response *types.StartGameResponse, err error)

	CloseGame(ctx context.Context, request *types.CloseGameRequest) (response *types.CloseGameResponse, err error)

	FindGame(ctx context.Context, request *types.FindGameRequest) (response *types.FindGameResponse, err error)

	FindGames(ctx context.Context, request *types.FindGamesRequest) (response *types.FindGamesResponse, err error)
}

type Service struct {
	*gameImpl
}

func NewService(
	mysql *xorm.Mysql,
	kubeclientset kubernetes.Interface,
	gamesclientset gamesclientset.Interface) GameService {
	return &Service{
		gameImpl: NewGame(mysql, kubeclientset, gamesclientset),
	}
}
