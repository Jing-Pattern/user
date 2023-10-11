package logic

import (
	"context"
	"fmt"
	"github.com/Jing-Pattern/user/rpc/internal/dao/query"

	"github.com/Jing-Pattern/user/rpc/internal/svc"
	"github.com/Jing-Pattern/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindUserLogic) FindUser(in *pb.UserReq) (*pb.ExistUser, error) {
	query := query.Use(l.svcCtx.Db).LoveUserInfo
	_, err := query.WithContext(context.Background()).Where(query.OpenID.Eq(in.Id)).First()
	if err != nil {
		fmt.Println("user1 err", err)
		return &pb.ExistUser{
			IsExist: false,
		}, err
	}

	return &pb.ExistUser{IsExist: true}, nil
}
