package user

import (
	"context"

	"kong-task/cmd/internal/svc"
	"kong-task/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) error {
	err := checkUserOperation(l.ctx, l.svcCtx, l.Logger, req.Organization)
	if err != nil {
		return err
	}

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.Uid, req.Project)
	if err != nil {
		return err
	}

	user.ActionType = req.ActionType
	user.IsAdmin = req.IsAdmin
	err = l.svcCtx.UserModel.Update(l.ctx, user)

	return err
}
