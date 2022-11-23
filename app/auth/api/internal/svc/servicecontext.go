package svc

import (
	"app/auth/api/internal/config"
	"app/tools/mysql"
	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	Mysql  *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Mysql:  mysql.GetDB(c.Mysql),
	}
}
