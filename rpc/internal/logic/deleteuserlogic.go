package logic

import (
	"context"

	"github.com/Jing-Pattern/user/rpc/internal/svc"
	"github.com/Jing-Pattern/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserLogic {
	return &DeleteUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DeleteUserLogic) DeleteUser(in *pb.UserReq) (*pb.UserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserResp{}, nil
}
