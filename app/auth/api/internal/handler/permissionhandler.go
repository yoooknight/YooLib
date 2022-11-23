package handler

import (
	"app/auth/api/internal/logic"
	"app/auth/api/internal/svc"
	"app/auth/api/internal/types"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

// PermissionListHandler 权限列表控制器
func PermissionListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req  types.PermissionListReq
			resp []types.PermissionListResp
			err  error
		)

		if err = httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		resp, err = logic.NewPermissionLogic(r.Context(), svcCtx).List()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
