package handler

import (
	"app/auth/api/internal/logic"
	"app/auth/api/internal/svc"
	"app/auth/api/internal/types"
	"app/tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req  types.LoginReq
			resp *types.LoginResp
			err  error
		)

		if err = httpx.Parse(r, &req); err != nil {
			response.Json(w, err.Error(), response.Fail)
			return
		}

		if resp, err = logic.NewLoginLogic(r.Context(), svcCtx).Login(req); err != nil {
			response.Json(w, err.Error(), response.InvalidParams)
			return
		}

		response.Json(w, resp)
	}
}
