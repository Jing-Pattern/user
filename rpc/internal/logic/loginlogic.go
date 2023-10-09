package logic

import (
	"context"
	"fmt"
	"github.com/Jing-Pattern/user/rpc/internal/dao/query"
	"strconv"

	"github.com/Jing-Pattern/user/rpc/internal/svc"
	"github.com/Jing-Pattern/user/rpc/pb/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *pb.LoginRequest) (*pb.LoginResponse, error) {
	// todo: add your logic here and delete this line
	u := query.Use(l.svcCtx.Db).LoveUserInfo

	code, _ := strconv.Atoi(in.Username)
	user, err := u.WithContext(l.ctx).Where(u.ID.Eq(int32(code))).First()
	if err != nil {
		return &pb.LoginResponse{}, err
	}
	fmt.Println(user)
	return &pb.LoginResponse{}, nil
}
