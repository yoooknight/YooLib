package logic

import (
	"app/auth/api/internal/svc"
	"app/auth/api/internal/types"
	"app/models"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type PermissionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPermissionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PermissionLogic {
	return &PermissionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PermissionLogic) List() (resp types.PermissionListResp, err error) {
	var tempResp []models.Permission

	// 这里处理业务逻辑
	var db = l.svcCtx.Mysql.Model(&models.Permission{})
	db.Scan(&tempResp)

	fmt.Printf("%+v", tempResp)
	fmt.Printf("%+v", resp)

	return
}
