package svc

import (
	"account/api/internal/config"
	"account/rpc/account"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config     config.Config
	AccountRpc account.Account
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:     c,
		AccountRpc: account.NewAccount(zrpc.MustNewClient(c.AccountRpc)),
	}
}
