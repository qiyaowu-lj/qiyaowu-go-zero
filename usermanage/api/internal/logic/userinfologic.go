package logic

import (
	"context"

	"Ningxi-Compose/usermanage/api/internal/svc"
	"Ningxi-Compose/usermanage/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserinfoRequest) (resp *types.UserinfoResponse, err error) {
	// todo: add your logic here and delete this line

	return
}