package logic

import (
	"context"
	"errors"
	"github.com/Jing-Pattern/user/rpc/internal/dao/model"
	"github.com/Jing-Pattern/user/rpc/internal/dao/query"
	"github.com/Jing-Pattern/user/rpc/internal/svc"
	"github.com/Jing-Pattern/user/rpc/pb"
	"github.com/zeromicro/go-zero/core/logx"
)

type BindLoverLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBindLoverLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BindLoverLogic {
	return &BindLoverLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *BindLoverLogic) BindLover(in *pb.LoverId) (*pb.UserResp, error) {
	q := query.Use(l.svcCtx.Db)
	userModel := q.LoveUserInfo
	relationModel := q.LoveLoverRelation

	var user, user2 *model.LoveUserInfo
	var userErr, user2Err error

	userCh := make(chan struct {
		user *model.LoveUserInfo
		err  error
	})
	user2Ch := make(chan struct {
		user *model.LoveUserInfo
		err  error
	})
	createRelationCh := make(chan error)

	go func() {
		user, userErr = userModel.WithContext(l.ctx).Where(userModel.OpenID.Eq(in.UserId)).First()
		userCh <- struct {
			user *model.LoveUserInfo
			err  error
		}{user, userErr}
	}()

	go func() {
		user2, user2Err = userModel.WithContext(l.ctx).Where(userModel.OpenID.Eq(in.LoverId)).First()
		user2Ch <- struct {
			user *model.LoveUserInfo
			err  error
		}{user2, user2Err}
	}()

	select {
	case userResult := <-userCh:
		if userResult.err != nil {
			return nil, errors.New("未找到用户")
		}
		if userResult.user.IsSingle == 2 {
			return nil, errors.New("对方已有情侣")
		}
		user = userResult.user
	case user2Result := <-user2Ch:
		if user2Result.err != nil {
			return nil, errors.New("未找到绑定用户")
		}
		user2 = user2Result.user
	}

	go func() {
		newRelation := model.LoveLoverRelation{User1ID: user.OpenID, User2ID: user2.OpenID}
		err := relationModel.WithContext(l.ctx).Create(&newRelation)
		createRelationCh <- err
	}()

	select {
	case err := <-createRelationCh:
		if err != nil {
			return nil, errors.New("创建用户关系失败")
		}
	default:
	}

	_, updateErr := userModel.WithContext(l.ctx).Where(userModel.OpenID.Eq(user.OpenID)).Or(userModel.OpenID.Eq(user2.OpenID)).Update(userModel.IsSingle, 2)
	if updateErr != nil {
		return nil, errors.New("更新情侣状态失败")
	}

	return &pb.UserResp{
		Message: "绑定成功",
	}, nil
}
