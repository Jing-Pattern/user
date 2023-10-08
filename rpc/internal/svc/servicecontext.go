package svc

import (
	"github.com/Jing-Pattern/user/rpc/internal/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := gorm.Open(mysql.Open(c.Mysql.DbUrl), &gorm.Config{})
	if err != nil {
		panic("DB初始化失败")

	}
	return &ServiceContext{
		Config: c,
		Db:     db,
	}
}
