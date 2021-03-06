// Code generated by goctl. DO NOT EDIT!
// Source: rpc.proto

package server

import (
	"context"

	"shopping/apps/cart/rpc/internal/logic"
	"shopping/apps/cart/rpc/internal/svc"
	"shopping/apps/cart/rpc/rpc"
)

type RpcServer struct {
	svcCtx *svc.ServiceContext
	rpc.UnimplementedRpcServer
}

func NewRpcServer(svcCtx *svc.ServiceContext) *RpcServer {
	return &RpcServer{
		svcCtx: svcCtx,
	}
}

func (s *RpcServer) Ping(ctx context.Context, in *rpc.Request) (*rpc.Response, error) {
	l := logic.NewPingLogic(ctx, s.svcCtx)
	return l.Ping(in)
}
