package user

import (
	"context"

	"kong-task/cmd/internal/svc"
	"kong-task/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUsersLogic {
	return &GetUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUsersLogic) GetUsers(req *types.GetUsersReq) (resp *types.GetUsersResp, err error) {
	err = checkUserOperation(l.ctx, l.svcCtx, l.Logger, req.Organization)
	if err != nil {
		return
	}

	users, err := l.svcCtx.UserModel.FindByProject(l.ctx, req.Project)
	if err != nil {
		return
	}

	resp = &types.GetUsersResp{
		Users: make([]types.UseInfo, len(users)),
	}

	for i, u := range users {
		resp.Users[i] = types.UseInfo{
			Uid:          u.Uid,
			Name:         u.Name,
			Organization: u.Organization,
			Project:      u.Project,
			IsAdmin:      u.IsAdmin,
			ActionType:   u.ActionType,
		}
	}

	return
}
