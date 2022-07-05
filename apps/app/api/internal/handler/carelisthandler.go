package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"shopping/apps/app/api/internal/logic"
	"shopping/apps/app/api/internal/svc"
	"shopping/apps/app/api/internal/types"
)

func CareListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CartListRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCareListLogic(r.Context(), svcCtx)
		resp, err := l.CareList(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
