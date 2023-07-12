package handler

import (
	"app/tools/response"
	"net/http"

	"app/book/internal/logic"
	"app/book/internal/svc"
	"app/book/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func BookHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			req  types.Request
			resp *types.Response
			err  error
		)

		if err = httpx.Parse(r, &req); err != nil {
			response.Json(w, response.Fail, err.Error())
			return
		}

		//l := logic.NewBookLogic(r.Context(), svcCtx)
		if resp, err = logic.NewBookLogic(r.Context(), svcCtx).Book(&req); err != nil {
			response.Json(w, response.Fail)
			return
		}

		response.Json(w, resp)
	}
}
