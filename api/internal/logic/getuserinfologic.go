package logic

import (
	"context"

	"account/api/internal/svc"
	"account/api/internal/types"
	"account/rpc/types/account"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.UserInfoReq) (resp *types.UserInfoRes, err error) {
	logx.Info("req.UserId==", req.UserId)
	res, err := l.svcCtx.AccountRpc.GetUserInfo(l.ctx, &account.UserInfoRequest{
		Userid: req.UserId,
	})

	if err != nil {
		return &types.UserInfoRes{
			Result: 2,
		}, err
	}
	return &types.UserInfoRes{
		Result:   1,
		UserId:   res.Info.Userid,
		Username: res.Info.Username,
		Sex:      res.Info.Sex,
		City:     res.Info.City,
		SchoolId: res.Info.Schoolid,
		Status:   int8(res.Info.Status),
		RegTime:  res.Info.RegTime,
		Phone:    res.Info.Phone,
		Email:    res.Info.Email,
	}, nil
}
