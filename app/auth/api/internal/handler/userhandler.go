package handler

import (
	"app/auth/api/internal/logic"
	"app/auth/api/internal/svc"
	"app/auth/api/internal/types"
	"app/models"
	"app/tools/response"
	"github.com/zeromicro/go-zero/rest/httpx"
	"net/http"
)

func UserInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req  types.UserInfoReq
			err  error
			resp models.User
		)

		if err = httpx.Parse(r, &req); err != nil {
			response.Json(w, err.Error(), response.Fail)
			return
		}

		if resp, err = logic.NewUserLogic(r.Context(), svcCtx).Info(req); err != nil {
			response.Json(w, err.Error(), response.Fail)
			return
		}

		response.Json(w, resp)
	}
}
