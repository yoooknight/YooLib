package logic

import (
	"app/auth/api/internal/svc"
	"app/auth/api/internal/types"
	"app/models"
	"context"
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

// List 获取权限列表
func (l *PermissionLogic) List() (resp []types.PermissionListResp, err error) {
	// 这里处理业务逻辑
	resp, _, err = l.getChildForMenu(0)
	return
}

// getChildForMenu 获取目录下的子菜单
func (l *PermissionLogic) getChildForMenu(menuId int64) (resp []types.PermissionListResp, isHasChildren bool, err error) {
	var (
		db       = l.svcCtx.Mysql.Model(&models.Permission{})
		tempList []models.Permission
	)

	db.Where("parent_id = ?", menuId).Scan(&tempList)
	if tempList != nil {
		isHasChildren = true
		// 如果存在，则表示有子集合
		for _, v := range tempList {
			// 获取子集
			tempChildren, tempIsHasChildren, err := l.getChildForMenu(v.ID)
			if err == nil {
				resp = append(resp, types.PermissionListResp{
					PermissionInfo: v,
					IsHasChild:     tempIsHasChildren,
					Children:       tempChildren,
				})
			}
		}
	} else {
		// 不存在，则无子集合
		isHasChildren = false
	}

	return
}
