package user

import (
	"context"

	"lebron/apps/app/api/internal/svc"
	"lebron/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserCollectionDelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserCollectionDelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserCollectionDelLogic {
	return &UserCollectionDelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserCollectionDelLogic) UserCollectionDel(req *types.UserCollectionDelReq) (resp *types.UserCollectionDelRes, err error) {
	// todo: add your logic here and delete this line

	return
}
