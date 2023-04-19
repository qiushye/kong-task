package user

import (
	"context"
	"fmt"

	"kong-task/cmd/internal/client"
	"kong-task/cmd/internal/model"
	"kong-task/cmd/internal/svc"
	"kong-task/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAddUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddUserLogic {
	return &AddUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddUserLogic) AddUser(req *types.UseInfo) error {

	err := checkUserOperation(l.ctx, l.svcCtx, l.Logger, req.Organization)
	if err != nil {
		return err
	}

	_, err = l.svcCtx.UserModel.FindOne(l.ctx, req.Uid, req.Project)
	if err == nil {
		return fmt.Errorf("user already exists")
	}

	err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		Uid:          req.Uid,
		Name:         req.Name,
		Project:      req.Project,
		Organization: req.Organization,
		IsAdmin:      req.IsAdmin,
		ActionType:   req.ActionType,
	})

	return err
}

func checkUserOperation(ctx context.Context, svcCtx *svc.ServiceContext, logger logx.Logger, organization string) (err error) {
	u, ok := client.FromContext(ctx)
	if !ok {
		return fmt.Errorf("auth failed")
	}

	defer func() {
		if err != nil {
			logger.Errorf("checkUserOperation(%+v) failed: %v", u, err)
		}
	}()

	operator, err := svcCtx.UserModel.FindByUidAndOrganization(ctx, u.Uid, organization)
	if err != nil {
		logger.Errorf("find operator %d %s failed: %v", u.Uid, organization, err)
		return fmt.Errorf("user not belong to the organization")
	}

	if !operator.IsAdmin {
		return fmt.Errorf("only the organization admin can add user")
	}

	return nil
}
