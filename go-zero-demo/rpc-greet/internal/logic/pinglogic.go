package logic

import (
	"context"

	"rpc-greet/internal/svc"
	"rpc-greet/transform"

	"github.com/tal-tech/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *transform.Request) (*transform.Response, error) {
	// todo: add your logic here and delete this line

	return &transform.Response{}, nil
}
