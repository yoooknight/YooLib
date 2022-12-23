package types

import "app/models"

type (
	PermissionListReq struct{}

	PermissionListResp struct {
		models.Permission
		HasChildren bool                 `json:"has_children"`
		Children    []PermissionListResp `json:"children"`
	}

	PermissionAddReq struct {
		ParentID      int64  `form:"parent_id"`                                                                 // 父级ID
		Name          string `form:"name" validate:"required,max=32" message:"请输入标题,标题不得超过32个字符"`               // 权限标题
		ServerRouter  string `form:"server_router" validate:"max=64" message:"路由不得超过64个字符"`                     // 后端路由
		WebRouter     string `form:"web_router" validate:"max=64" message:"路由不得超过64个字符"`                        // 前端路由
		RequestMethod string `form:"request_method" validate:"required,max=10" message:"请选择请求方式,请求方式不得超过10个字符"` // 请求方式 get post
		Status        int64  `form:"status" validate:"oneof=1 2" message:"状态参数格式错误"`                            // 状态 1启用 2禁用
	}
)
