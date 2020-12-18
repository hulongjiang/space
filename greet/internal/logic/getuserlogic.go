package logic

import (
	"context"
	"fmt"
	"space/user/user"

	"space/greet/internal/svc"
	"space/greet/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) GetUserLogic {
	return GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req types.GetUserRequest) (*types.Response, error) {
	rsp := &types.Response{}

	fmt.Println(req.Id)

	userRsp, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.GetUserReq{Id: req.Id})
	if err != nil { // code error
		return rsp, err
	}

	rsp.Message = userRsp.Name

	return rsp, nil
}
