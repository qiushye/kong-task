package user

import (
	"context"

	"kong-task/cmd/internal/svc"
	"kong-task/cmd/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/mongo"
)

type RemoveUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveUserLogic {
	return &RemoveUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveUserLogic) RemoveUser(req *types.RemoveUserReq) error {
	err := checkUserOperation(l.ctx, l.svcCtx, l.Logger, req.Organization)
	if err != nil {
		return err
	}

	_, err = l.svcCtx.UserModel.FindOne(l.ctx, req.Uid, req.Project)
	if err != nil {
		if err.Error() == mongo.ErrNoDocuments.Error() {
			return nil
		}
		return err
	}

	err = l.svcCtx.UserModel.Delete(l.ctx, req.Uid, req.Project)

	return err

}
