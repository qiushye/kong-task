package client

import "context"

var userInfoKey = "user-info"

type UserInfo struct {
	Uid uint32 `json:"Uid"`
}

const (
	ActionNone uint = iota
	ActionRO
	ActionRW
)

func NewContext(ctx context.Context, u UserInfo) context.Context {
	return context.WithValue(ctx, userInfoKey, u)
}

func FromContext(ctx context.Context) (UserInfo, bool) {
	u, ok := ctx.Value(userInfoKey).(UserInfo)
	return u, ok
}
