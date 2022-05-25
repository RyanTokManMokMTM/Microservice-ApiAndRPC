// Code generated by goctl. DO NOT EDIT!
// Source: product.proto

package user

import (
	"context"

	"mall/service/user/rpc/type/user"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	LoginRequest     = user.LoginRequest
	LoginResponse    = user.LoginResponse
	RegisterRequest  = user.RegisterRequest
	RegisterResponse = user.RegisterResponse
	UserInfoRequest  = user.UserInfoRequest
	UserInfoResponse = user.UserInfoResponse

	USer interface {
		Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error)
		Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error)
		UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error)
	}

	defaultUSer struct {
		cli zrpc.Client
	}
)

func NewUSer(cli zrpc.Client) USer {
	return &defaultUSer{
		cli: cli,
	}
}

func (m *defaultUSer) Login(ctx context.Context, in *LoginRequest, opts ...grpc.CallOption) (*LoginResponse, error) {
	client := user.NewUSerClient(m.cli.Conn())
	return client.Login(ctx, in, opts...)
}

func (m *defaultUSer) Register(ctx context.Context, in *RegisterRequest, opts ...grpc.CallOption) (*RegisterResponse, error) {
	client := user.NewUSerClient(m.cli.Conn())
	return client.Register(ctx, in, opts...)
}

func (m *defaultUSer) UserInfo(ctx context.Context, in *UserInfoRequest, opts ...grpc.CallOption) (*UserInfoResponse, error) {
	client := user.NewUSerClient(m.cli.Conn())
	return client.UserInfo(ctx, in, opts...)
}
