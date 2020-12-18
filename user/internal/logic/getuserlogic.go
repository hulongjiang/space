package logic

import (
	"context"

	"space/user/internal/svc"
	"space/user/user"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserLogic) GetUser(in *user.GetUserReq) (*user.GetUserRsp, error) {
	// todo: add your logic here and delete this line
	rsp := &user.GetUserRsp{}
	rsp.Id = in.Id
	rsp.Name = "胡龙江"
	rsp.Pwd = "123456"

	return rsp, nil
}
