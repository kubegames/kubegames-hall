package room

import (
	"context"

	"github.com/kubegames/kubegames-hall/app/service/room/types"
	"github.com/kubegames/kubegames-hall/internal/pkg/nsq"
	"github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
)

type RoomService interface {
	AddGameRoom(ctx context.Context, request *types.AddGameRoomRequest) (response *types.AddGameRoomResponse, err error)

	AddGameRooms(ctx context.Context, request *types.AddGameRoomsRequest) (response *types.AddGameRoomsResponse, err error)

	CloseGameRoom(ctx context.Context, request *types.CloseGameRoomRequest) (response *types.CloseGameRoomResponse, err error)

	FindGameRoom(ctx context.Context, request *types.FindGameRoomRequest) (response *types.FindGameRoomResponse, err error)

	FindGameRooms(ctx context.Context, request *types.FindGameRoomsRequest) (response *types.FindGameRoomsResponse, err error)

	OpenGameRoom(ctx context.Context, request *types.OpenGameRoomRequest) (response *types.OpenGameRoomResponse, err error)

	AddGameRoomPools(ctx context.Context, request *types.AddGameRoomPoolsRequest) (response *types.AddGameRoomPoolsResponse, err error)

	BindGameRoomPool(ctx context.Context, request *types.BindGameRoomPoolRequest) (response *types.BindGameRoomPoolResponse, err error)

	UnboundGameRoomPool(ctx context.Context, request *types.UnboundGameRoomPoolRequest) (response *types.UnboundGameRoomPoolResponse, err error)

	DeleteGameRoomPool(ctx context.Context, request *types.DeleteGameRoomPoolRequest) (response *types.DeleteGameRoomPoolResponse, err error)

	FindGameRoomPool(ctx context.Context, request *types.FindGameRoomPoolRequest) (response *types.FindGameRoomPoolResponse, err error)

	FindGameRoomPools(ctx context.Context, request *types.FindGameRoomPoolsRequest) (response *types.FindGameRoomPoolsResponse, err error)
}

type Service struct {
	*roomImpl
}

func NewService(mysql *xorm.Mysql, nsq nsq.Nsq) RoomService {
	return &Service{
		roomImpl: NewRoom(mysql, nsq),
	}
}
