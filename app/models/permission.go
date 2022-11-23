package models

type Permission struct {
	BaseModel
	ParentID      int64  `json:"parent_id"`                       //父级ID
	Name          string `json:"name"`                            //标题
	WebRouter     string `json:"web_router"`                      //前端路由
	ServerRouter  string `json:"server_router"`                   //服务端路由
	RequestMethod string `json:"request_method"`                  //请求方式 get post put delete
	Status        int64  `json:"status"`                          //状态 1启用 2禁用
	DeleteDate    string `json:"delete_date" gorm:"default:NULL"` //删除时间
}
