package logic

import (
	"context"

	"app/auth/rpc/internal/svc"
	"app/auth/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAuthByRoleIDLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAuthByRoleIDLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAuthByRoleIDLogic {
	return &UpdateAuthByRoleIDLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAuthByRoleIDLogic) UpdateAuthByRoleID(in *pb.UpdateAuthByRoleIDReq) (*pb.Resp, error) {
	// todo: add your logic here and delete this line

	return &pb.Resp{}, nil
}
