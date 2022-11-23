package config

import (
	"app/tools/mysql"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Mysql mysql.Config
}
