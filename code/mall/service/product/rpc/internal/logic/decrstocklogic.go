package logic

import (
	"context"
	"database/sql"
	"github.com/dtm-labs/dtmcli"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/dtm-labs/dtmgrpc"
	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DecrStockLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDecrStockLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DecrStockLogic {
	return &DecrStockLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// DTM function
func (l *DecrStockLogic) DecrStock(in *product.DetailRequest) (*product.DecrStockResponse, error) {
	// todo: add your logic here and delete this line

	//getting Raw DB
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	//getting DTM Barrier Object
	dtmBarrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	//Transaction with Barrier
	err = dtmBarrier.CallWithDB(db, func(tx *sql.Tx) error {
		//update stock
		res, err := l.svcCtx.ProductModel.TxAdjustStock(tx, in.Id, -1)
		if err != nil {
			return err
		}

		affected, err := res.RowsAffected()
		//transaction failed
		if err == nil && affected == 0 {
			return dtmcli.ErrFailure
		}

		return err
	})

	//sock < 1
	if err == dtmcli.ErrFailure {
		//rollback
		return nil, status.Error(codes.Aborted, dtmcli.ResultFailure)
	}

	if err != nil {
		return nil, err
	}
	return &product.DecrStockResponse{}, nil
}
