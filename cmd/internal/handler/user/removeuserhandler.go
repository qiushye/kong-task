package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"kong-task/cmd/internal/logic/user"
	"kong-task/cmd/internal/svc"
	"kong-task/cmd/internal/types"
)

func RemoveUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.RemoveUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := user.NewRemoveUserLogic(r.Context(), svcCtx)
		err := l.RemoveUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
