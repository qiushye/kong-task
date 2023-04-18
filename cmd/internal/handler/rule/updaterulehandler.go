package rule

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"kong-task/cmd/internal/logic/rule"
	"kong-task/cmd/internal/svc"
	"kong-task/cmd/internal/types"
)

func UpdateRuleHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateRuleReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := rule.NewUpdateRuleLogic(r.Context(), svcCtx)
		err := l.UpdateRule(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
