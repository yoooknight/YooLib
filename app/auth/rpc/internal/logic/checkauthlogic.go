package logic

import (
	"context"

	"app/auth/rpc/internal/svc"
	"app/auth/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckAuthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckAuthLogic {
	return &CheckAuthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CheckAuthLogic) CheckAuth(in *pb.CheckAuthReq) (*pb.Resp, error) {
	// todo: add your logic here and delete this line

	return &pb.Resp{}, nil
}
