package logic

import (
	"context"

	"app/auth/rpc/internal/svc"
	"app/auth/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateAllAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateAllAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateAllAuthLogic {
	return &UpdateAllAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateAllAuthLogic) UpdateAllAuth(in *pb.Empty) (*pb.Resp, error) {
	// todo: add your logic here and delete this line

	return &pb.Resp{}, nil
}
