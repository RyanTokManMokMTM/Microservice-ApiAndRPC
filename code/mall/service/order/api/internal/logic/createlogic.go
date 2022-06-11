package logic

import (
	"context"
	"github.com/dtm-labs/dtmgrpc"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/types/order"
	"mall/service/product/rpc/types/product"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line

	//getting orderRPC buildTarget
	orderRPCBuildServer, err := l.svcCtx.Config.OrderRPC.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "order created failed with unknown issues")
	}

	//getting productRPC buildServer
	productRPCBuildServer, err := l.svcCtx.Config.ProductRPC.BuildTarget()
	if err != nil {
		return nil, status.Error(100, "order created failed with unknown issues")
	}

	dtmServer := "etcd://etcd:2379/dtmservice"

	//creating an gid
	gid := dtmgrpc.MustGenGid(dtmServer)

	//saga Protocol
	saga := dtmgrpc.NewSagaGrpc(dtmServer, gid).
		Add(orderRPCBuildServer+"/order.Order/Create", orderRPCBuildServer+"/order.Order/CreateRevert", &order.CreateRequest{
			Uid:    req.Uid,
			Pid:    req.Pid,
			Amount: req.Amount,
			Status: 0,
		}).Add(productRPCBuildServer+"/product.Product/DecrStock", productRPCBuildServer+"/product.Product/Create/DecrStockRevert", &product.DecrStockRequest{
		Id:  req.Pid,
		Num: 1,
	})

	err = saga.Submit()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &types.CreateResponse{}, nil
}
