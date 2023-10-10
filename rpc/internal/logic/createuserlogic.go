package logic

import (
	"context"

	"github.com/Jing-Pattern/user/rpc/internal/svc"
	"github.com/Jing-Pattern/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *pb.UserReq) (*pb.UserResp, error) {
	// todo: add your logic here and delete this line

	return &pb.UserResp{}, nil
}
