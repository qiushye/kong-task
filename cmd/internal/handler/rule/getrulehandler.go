package rule

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"kong-task/cmd/internal/logic/rule"
	"kong-task/cmd/internal/svc"
	"kong-task/cmd/internal/types"
)

func GetRuleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetRuleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := rule.NewGetRuleLogic(r.Context(), svcCtx)
		resp, err := l.GetRule(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
