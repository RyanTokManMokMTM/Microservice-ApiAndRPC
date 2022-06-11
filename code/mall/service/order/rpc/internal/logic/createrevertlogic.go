package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dtm-labs/dtmgrpc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"google.golang.org/grpc/status"
	"mall/service/user/rpc/type/user"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateRevertLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateRevertLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateRevertLogic {
	return &CreateRevertLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// revert method
func (l *CreateRevertLogic) CreateRevert(in *order.CreateRequest) (*order.CreateResponse, error) {
	// todo: add your logic here and delete this line

	//getting raw db
	db, err := sqlx.NewMysql(l.svcCtx.Config.Mysql.DataSource).RawDB()
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	//getting barrier
	barrier, err := dtmgrpc.BarrierFromGrpc(l.ctx)
	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	err = barrier.CallWithDB(db, func(tx *sql.Tx) error {

		//check user
		_, err := l.svcCtx.UserRPC.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid})
		if err != nil {
			return err
		}

		//check order
		res, err := l.svcCtx.OrderModel.FindOneByUid(in.Uid)
		if err != nil {
			return fmt.Errorf("order not exist")
		}

		res.Status = 9 //order is not available
		err = l.svcCtx.OrderModel.TxUpdate(tx, res)
		if err != nil {
			return fmt.Errorf("update order failed")
		}
		return nil
	})

	if err != nil {
		return nil, status.Error(500, err.Error())
	}

	return &order.CreateResponse{}, nil
}
