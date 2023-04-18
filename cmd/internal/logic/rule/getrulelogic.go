package rule

import (
	"context"

	"kong-task/cmd/internal/svc"
	"kong-task/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetRuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetRuleLogic {
	return &GetRuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetRuleLogic) GetRule(req *types.GetRuleReq) (resp *types.GetRuleResp, err error) {
	// todo: add your logic here and delete this line

	return
}
