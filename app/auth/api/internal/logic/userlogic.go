package logic

import (
	"app/auth/api/internal/svc"
	"app/auth/api/internal/types"
	"app/models"
	"context"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Info 获取用户信息
func (u *UserLogic) Info(req types.UserInfoReq) (resp types.UserInfoResp, err error) {
	var (
		userInfo models.User
		menu     []types.PermissionListResp
	)

	if err = u.svcCtx.Mysql.Model(models.User{}).Where("id=?", u.ctx.Value("uid")).First(&userInfo).Error; err != nil {
		return
	} else {

	}

	// 获取权限
	menu, err = NewPermissionLogic(u.ctx, u.svcCtx).List()

	// 拼接返回值
	resp = types.UserInfoResp{
		User: userInfo,
		Menu: menu,
	}

	return
}

// Login 用户登录
func (u *UserLogic) Login(req types.LoginReq) (resp *types.LoginResp, err error) {
	var userInfo models.User

	resp = new(types.LoginResp)

	// 验证数据库是否有该用户
	if err = u.svcCtx.Mysql.Model(models.User{}).Where("user_name=? and status=1", req.UserName).First(&userInfo).Error; err != nil {
		return
	}

	if userInfo.Password != req.Password {
		err = errors.New("密码错误")
		return
	}

	resp.Token, err = u.getToken(time.Now().Unix(), u.svcCtx.Config.Auth.AccessSecret, map[string]interface{}{"uid": 1, "username": "root"}, u.svcCtx.Config.Auth.AccessExpire)
	return
}

// getToken 获取token值
func (u *UserLogic) getToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat

	for k, v := range payloads {
		claims[k] = v
	}

	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims

	return token.SignedString([]byte(secretKey))
}
