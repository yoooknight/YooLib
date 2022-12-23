package config

import (
	"app/tools/mysql"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql mysql.Config
	Auth  JwtAuth
}

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}
