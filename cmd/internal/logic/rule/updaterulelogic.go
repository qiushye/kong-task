package rule

import (
	"context"
	"fmt"

	"kong-task/cmd/internal/client"
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
	u, ok := client.FromContext(l.ctx)
	if !ok {
		return fmt.Errorf("auth failed")
	}

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, u.Uid, req.Project)
	if err != nil {
		return fmt.Errorf("invalid user")
	}

	if user.ActionType < client.ActionRW {
		return fmt.Errorf("you are unable to update rule")
	}

	rule, err := l.svcCtx.RuleModel.FindOne(l.ctx, req.Name)
	if err != nil {
		l.Logger.Errorf("find rule(%v) failed: %v", req.Name, err)
		return err
	}

	if rule.Project != req.Project {
		return fmt.Errorf("the rule doesn't belong to the project %s", req.Project)
	}

	rule.Content = req.Content
	err = l.svcCtx.RuleModel.Update(l.ctx, rule)
	if err != nil {
		l.Logger.Errorf("insert rule(%+v) failed: %v", *req, err)
		return err
	}

	return nil
}
