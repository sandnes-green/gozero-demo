package logic

import (
	"context"

	"account/rpc/internal/svc"
	"account/rpc/types/account"

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

func (l *GetUserInfoLogic) GetUserInfo(in *account.UserInfoRequest) (*account.UserInfoResponse, error) {
	logx.Info("in===", *in)
	one, err := l.svcCtx.Model.GetUserInfo(l.ctx, in.Userid)
	if err != nil {
		return nil, err
	}
	logx.Info("one===", *one)
	return &account.UserInfoResponse{
		Result: 1,
		Info: &account.UserInfo{
			Userid:   one.FldUserid,
			Username: one.FldName,
			Sex:      int32(one.FldSex),
			Mood:     one.FldMood,
			City:     one.FldCity,
			Schoolid: one.FldSchoolId,
			Status:   int32(one.FldStatus),
			RegTime:  one.FldRegistTime.Format("2006-01-02 15:04:05"),
			Phone:    one.FldBindMobile,
			Email:    one.FldBindEmail,
		},
	}, nil
}
