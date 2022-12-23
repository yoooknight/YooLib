package logic

import (
	"app/auth/api/internal/svc"
	"app/auth/api/internal/types"
	"app/models"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
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
func (u *UserLogic) Info(req types.UserInfoReq) (resp models.User, err error) {
	err = u.svcCtx.Mysql.Model(models.User{}).First(&resp).Error
	return
}
