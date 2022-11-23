package mysql

import "C"
import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// Config 数据库配置信息
type Config struct {
	Username                  string // 用户名
	Password                  string // 密码
	Host                      string // 链接地址
	Port                      string // 端口
	Database                  string // 数据库名
	Charset                   string // 字符集
	ParseTime                 string // 是否解析时间
	Loc                       string // 时区默认 Local
	TablePrefix               string // 表前缀
	SingularTable             bool   // 是否使用单数表名 true 是 false 否
	PrepareStmt               bool   //  在执行任何 SQL 时都会创建一个 prepared statement 并将其缓存，以提高后续的效率
	ConnMaxLifetime           int    // 设置了连接可复用的最大时间
	MaxIdleConn               int    // 连接池里面的连接最大存活时长(秒)
	ConnMaxIdleTime           int    //连接池里面的连接最大空闲时长(秒)
	MaxOpenConn               int    // 设置打开数据库连接的最大数量
	IgnoreRecordNotFoundError bool   // 忽略ErrRecordNotFound（记录未找到）错误
}

func GetDB(c Config) (DB *gorm.DB) {
	var err error

	mysqlConfig := mysql.Config{
		DSN: fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=%s&loc=%s",
			c.Username, c.Password, c.Host, c.Port, c.Database, c.Charset, c.ParseTime, c.Loc),
	}

	DB, err = gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   c.TablePrefix,   //  表名前缀，`User` 的表名应该是 `t_users`
			SingularTable: c.SingularTable, //  使用单数表名，启用该选项，此时，`User` 的表名应该是 `t_user`
		},
	})
	if err != nil {
		panic(err.Error())
	}
	return
}
