package rule

import (
	"context"

	"kong-task/cmd/internal/svc"
	"kong-task/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateRuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRuleLogic {
	return &UpdateRuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateRuleLogic) UpdateRule(req *types.UpdateRuleReq) error {
	// todo: add your logic here and delete this line

	return nil
}
