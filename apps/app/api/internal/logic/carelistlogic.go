package logic

import (
	"context"

	"shopping/apps/app/api/internal/svc"
	"shopping/apps/app/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CareListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCareListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CareListLogic {
	return &CareListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CareListLogic) CareList(req *types.CartListRequest) (resp *types.CartListResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
