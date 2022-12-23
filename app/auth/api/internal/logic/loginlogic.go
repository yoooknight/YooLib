package logic

import (
	"app/auth/api/internal/svc"
	"app/auth/api/internal/types"
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"time"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Login 用户登录
func (l *LoginLogic) Login(req types.LoginReq) (resp *types.LoginResp, err error) {
	resp = new(types.LoginResp)
	resp.Token, err = l.getToken(time.Now().Unix(), l.svcCtx.Config.Auth.AccessSecret, map[string]interface{}{"uid": 1, "username": "root"}, l.svcCtx.Config.Auth.AccessExpire)
	return
}

// getToken 获取token值
func (l *LoginLogic) getToken(iat int64, secretKey string, payloads map[string]interface{}, seconds int64) (string, error) {
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
