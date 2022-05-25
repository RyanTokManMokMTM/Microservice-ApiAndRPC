// Code generated by goctl. DO NOT EDIT!
// Source: user.proto

package server

import (
	"context"

	"mall/service/user/rpc/internal/logic"
	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/type/user"
)

type USerServer struct {
	svcCtx *svc.ServiceContext
	user.UnimplementedUSerServer
}

func NewUSerServer(svcCtx *svc.ServiceContext) *USerServer {
	return &USerServer{
		svcCtx: svcCtx,
	}
}

func (s *USerServer) Login(ctx context.Context, in *user.LoginRequest) (*user.LoginResponse, error) {
	l := logic.NewLoginLogic(ctx, s.svcCtx)
	return l.Login(in)
}

func (s *USerServer) Register(ctx context.Context, in *user.RegisterRequest) (*user.RegisterResponse, error) {
	l := logic.NewRegisterLogic(ctx, s.svcCtx)
	return l.Register(in)
}

func (s *USerServer) UserInfo(ctx context.Context, in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	l := logic.NewUserInfoLogic(ctx, s.svcCtx)
	return l.UserInfo(in)
}