package middleware

import (
	"fmt"
	"kong-task/cmd/internal/client"
	"net/http"
	"net/url"
	"strconv"
)

type AuthMiddleware struct {
}

func NewAuthMiddleware() *AuthMiddleware {
	return &AuthMiddleware{}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		auth := r.Header.Get("Auth")
		userInfo, err := parseToken(auth)
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

// 这里的鉴权实际是明文，op-uid=*** 格式
func parseToken(auth string) (*client.UserInfo, error) {

	var (
		err error
	)

	values, err := url.ParseQuery(auth)
	if err != nil {
		return nil, fmt.Errorf("parse query failed: %s", err.Error())
	}
	opUid, err := strconv.ParseUint(values.Get("op_uid"), 10, 64)
	if err != nil {
		return nil, fmt.Errorf("parse op_uid failed: %s", err.Error())
	}

	u := &client.UserInfo{
		Uid: uint32(opUid),
	}
	return u, nil
}
