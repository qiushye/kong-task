package rule

import (
	"context"
	"fmt"

	"kong-task/cmd/internal/client"
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
	u, ok := client.FromContext(l.ctx)
	if !ok {
		return nil, fmt.Errorf("auth failed")
	}

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, u.Uid, req.Project)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	if user.ActionType < client.ActionRO {
		return nil, fmt.Errorf("you are unable to query rule")
	}

	rule, err := l.svcCtx.RuleModel.FindOne(l.ctx, req.Name)
	if err != nil {
		l.Logger.Errorf("insert rule(%+v) failed: %v", *req, err)
		return
	}

	if rule.Project != req.Project {
		return nil, fmt.Errorf("rule name doesn't belong to the project")
	}

	return &types.GetRuleResp{
		Id: rule.RuleId(),
		LintingRule: types.LintingRule{
			Name:    rule.Name,
			Project: rule.Project,
			Content: rule.Content,
		},
	}, nil
}
