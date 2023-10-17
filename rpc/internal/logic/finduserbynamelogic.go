package logic

import (
	"context"
	"github.com/Jing-Pattern/user/rpc/internal/dao/query"
	"github.com/jinzhu/copier"

	"github.com/Jing-Pattern/user/rpc/internal/svc"
	"github.com/Jing-Pattern/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserByNameLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserByNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserByNameLogic {
	return &FindUserByNameLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserByNameLogic) FindUserByName(in *pb.UserByNameReq) (*pb.UserInfo, error) {
	// todo: add your logic here and delete this line
	que := query.Use(l.svcCtx.Db).LoveUserInfo
	user := new(pb.UserInfo)
	userModel, err := que.WithContext(l.ctx).Where(que.Name.Eq(in.Name)).First()
	if err != nil {
		l.Logger.Info("未查到该用户")
	}
	err = copier.Copy(&user, userModel)
	if err != nil {
		return nil, err
	}
	return user, nil
}
