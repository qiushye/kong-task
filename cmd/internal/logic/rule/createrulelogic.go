package rule

import (
	"context"
	"fmt"

	"kong-task/cmd/internal/client"
	"kong-task/cmd/internal/model"
	"kong-task/cmd/internal/svc"
	"kong-task/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRuleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateRuleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRuleLogic {
	return &CreateRuleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateRuleLogic) CreateRule(req *types.CreateRuleReq) (resp *types.CreateRuleResp, err error) {
	u, ok := client.FromContext(l.ctx)
	if !ok {
		return nil, fmt.Errorf("auth failed")
	}

	if u.ActionType < client.ActionRW {
		return nil, fmt.Errorf("you are unable to create rule")
	}

	err = l.svcCtx.RuleModel.Insert(l.ctx, &model.Rule{
		Name:    req.Name,
		Project: req.Project,
		Content: req.Content,
	})
	if err != nil {
		l.Logger.Errorf("insert rule(%+v) failed: %v", *req, err)
		return
	}

	return
}
