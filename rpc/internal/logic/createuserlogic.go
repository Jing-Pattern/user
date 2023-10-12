package logic

import (
	"context"
	"github.com/Jing-Pattern/user/rpc/internal/dao/model"
	"github.com/Jing-Pattern/user/rpc/internal/dao/query"
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

func (l *CreateUserLogic) CreateUser(in *pb.CreateUserReq) (*pb.UserResp, error) {
	// todo: add your logic here and delete this line
	query := query.Use(l.svcCtx.Db).LoveUserInfo

	user := model.LoveUserInfo{Name: "Modi", OpenID: in.OpenId, Nickname: "请设置昵称"}
	err := query.WithContext(l.ctx).Create(&user)
	if err != nil {
		return &pb.UserResp{Message: "用户创建失败"}, err
	}
	return &pb.UserResp{Message: "用户创建成功"}, nil
}
