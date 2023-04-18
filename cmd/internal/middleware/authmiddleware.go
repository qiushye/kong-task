package middleware

import (
	"fmt"
	"kong-task/cmd/internal/client"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token := r.Header.Get("token")
		userInfo, err := parseToken(token)
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte(err.Error()))
			return
		}

		ctx := client.NewContext(r.Context(), *userInfo)

		// Passthrough to next handler if need
		next(w, r.WithContext(ctx))
	}
}

// 这里写死检查方法，暂时没有其他服务可以提高token校验
func parseToken(token string) (*client.UserInfo, error) {

	var (
		userInfo *client.UserInfo
		err      error
	)
	if strings.HasPrefix(token, "0") {

		userInfo = &client.UserInfo{
			Uid:          123,
			Name:         "Eric",
			Organization: "Kong",
			ActionType:   0,
		}
	} else if strings.HasPrefix(token, "1") {

		userInfo = &client.UserInfo{
			Uid:          456,
			Name:         "Tom",
			Organization: "Kong",
			ActionType:   1,
		}
	} else {

		err = fmt.Errorf("invalid token")
	}

	return userInfo, err
}
