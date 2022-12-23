package models

type User struct {
	ID       int64  `json:"id"`        // 用户自增id
	UserName string `json:"user_name"` // 用户名称
	Password string `json:"password"`  // 用户密码
	Status   int64  `json:"status"`    // 用户状态：1-启动 2-禁用
}
