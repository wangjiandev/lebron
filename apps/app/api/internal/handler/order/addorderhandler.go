package order

import (
	"net/http"

	"lebron/apps/app/api/internal/logic/order"
	"lebron/apps/app/api/internal/svc"
	"lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func AddOrderHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.OrderAddReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := order.NewAddOrderLogic(r.Context(), svcCtx)
		resp, err := l.AddOrder(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
