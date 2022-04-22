package platform

import (
	"context"

	"github.com/kubegames/kubegames-hall/app/service/platform/types"
	"github.com/kubegames/kubegames-hall/internal/pkg/nats"
	"github.com/kubegames/kubegames-hall/internal/pkg/nsq"
	"github.com/kubegames/kubegames-hall/internal/pkg/orm/xorm"
)

type GameService interface {
	ApplyRobot(ctx context.Context, request *types.ApplyRobotRequest) (response *types.ApplyRobotResponse, err error)

	PlayerLeaveGame(ctx context.Context, request *types.PlayerLeaveGameRequest) (response *types.PlayerLeaveGameResponse, err error)

	Broadcast(ctx context.Context, request *types.BroadcastRequest) (response *types.BroadcastResponse, err error)

	AcquireRoom(ctx context.Context, request *types.AcquireRoomRequest) (response *types.AcquireRoomResponse, err error)

	FindRooms(ctx context.Context, request *types.FindRoomsRequest) (response *types.FindRoomsResponse, err error)

	ReleaseRoom(ctx context.Context, request *types.ReleaseRoomRequest) (response *types.ReleaseRoomResponse, err error)

	RunningRoom(ctx context.Context, request *types.RunningRoomRequest) (response *types.RunningRoomResponse, err error)

	UploadPlayerRecord(ctx context.Context, request *types.UploadPlayerRecordRequest) (response *types.UploadPlayerRecordResponse, err error)

	UploadPlayerSettleInfo(ctx context.Context, request *types.UploadPlayerSettleInfoRequest) (response *types.UploadPlayerSettleInfoResponse, err error)

	GetRoomPool(ctx context.Context, request *types.GetRoomPoolRequest) (response *types.GetRoomPoolResponse, err error)
}

type Service struct {
	*platformImpl
}

func NewService(mysql *xorm.Mysql, nats nats.Nats, nsq nsq.Nsq) GameService {
	return &Service{
		platformImpl: NewPlatform(mysql, nats, nsq),
	}
}
