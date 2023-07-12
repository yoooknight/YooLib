package logic

import (
	"app/auth/rpc/authrpc"
	"app/book/internal/svc"
	"app/book/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type BookLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BookLogic {
	return &BookLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BookLogic) Book(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line'
	// 使用user rpc
	_, err = l.svcCtx.AuthRpc.CheckAuth(l.ctx, &authrpc.CheckAuthReq{
		UID:    "1",
		URL:    "/root/dn",
		Method: "get",
	})

	//resp.Message = "hahahah"

	return
}
