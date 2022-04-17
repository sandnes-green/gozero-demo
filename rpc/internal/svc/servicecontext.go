package svc

import (
	"account/rpc/internal/config"

	"account/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Model  model.TblAccountModel // 手动代码
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Model:  model.NewTblAccountModel(sqlx.NewMysql(c.DataSource), c.Cache), // 手动代码
	}
}
