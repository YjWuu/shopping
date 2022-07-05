package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"shopping/apps/app/api/internal/logic"
	"shopping/apps/app/api/internal/svc"
)

func FlashSaleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewFlashSaleLogic(r.Context(), svcCtx)
		resp, err := l.FlashSale()
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
