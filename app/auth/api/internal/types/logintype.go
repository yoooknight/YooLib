package types

type (
	// LoginReq 用户登录请求数据
	LoginReq struct {
		UserName string `json:"user_name" validate:"required, max=32" message:"请输入用户名,用户名不超过32个字符"` // 用户名
		Password string `json:"password" validate:"required, max=32" message:"请输入用户密码，密码不超过32个字符"`  // 用户密码
	}

	// LoginResp 用户登录返回数据
	LoginResp struct {
		Token string `form:"token"` // 用户token值
	}
)
