package logic

import (
	"context"
	"github.com/Jing-Pattern/user/rpc/internal/dao/query"
	"github.com/Jing-Pattern/user/rpc/internal/svc"
	"github.com/Jing-Pattern/user/rpc/pb"
	"github.com/jinzhu/copier"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *pb.UserReq) (*pb.UserInfo, error) {
	// todo: add your logic here and delete this line
	query := query.Use(l.svcCtx.Db).LoveUserInfo
	user := new(pb.UserInfo)
	model, err := query.WithContext(context.Background()).Where(query.OpenID.Eq(in.Id)).First()
	if err != nil {
		return &pb.UserInfo{}, err
	}
	copier.Copy(&user, model)
	return user, nil
}
