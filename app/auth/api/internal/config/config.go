package config

import (
	"app/tools/mysql"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql        mysql.Config
	Auth         JwtAuth
	SuperManager SuperManager
}

type JwtAuth struct {
	AccessSecret string
	AccessExpire int64
}

type SuperManager struct {
	Username string
}
