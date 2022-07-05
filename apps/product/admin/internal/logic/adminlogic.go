package logic

import (
	"context"

	"shopping/apps/product/admin/internal/svc"
	"shopping/apps/product/admin/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AdminLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAdminLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AdminLogic {
	return &AdminLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AdminLogic) Admin(req *types.Request) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
