package rule

import (
	"context"
	"fmt"

	"kong-task/cmd/internal/client"
	"kong-task/cmd/internal/model"
	"kong-task/cmd/internal/svc"
	"kong-task/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, u.Uid, req.Project)
	if err != nil {
		return nil, fmt.Errorf("invalid user")
	}

	if user.ActionType < client.ActionRW {
		return nil, fmt.Errorf("you are unable to create rule")
	}

	_, err = l.svcCtx.RuleModel.FindByProject(l.ctx, req.Project)
	if err == nil {
		return nil, fmt.Errorf("project %s already has rule", req.Project)
	}

	id := primitive.NewObjectID()
	err = l.svcCtx.RuleModel.Insert(l.ctx, &model.Rule{
		Id:      id,
		Name:    req.Name,
		Project: req.Project,
		Content: req.Content,
	})
	if err != nil {
		l.Logger.Errorf("insert rule(%+v) failed: %v", *req, err)
		return
	}

	return &types.CreateRuleResp{Id: id.Hex()}, nil
}
