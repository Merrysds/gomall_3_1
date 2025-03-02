package service

import (
	"context"

	"github.com/cloudwego/gomall07/app/cart/biz/dal/mysql"
	"github.com/cloudwego/gomall07/app/cart/biz/model"
	cart "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type EmptyCartService struct {
	ctx context.Context
} // NewEmptyCartService new EmptyCartService
func NewEmptyCartService(ctx context.Context) *EmptyCartService {
	return &EmptyCartService{ctx: ctx}
}

// Run create note info
func (s *EmptyCartService) Run(req *cart.EmptyCartReq) (resp *cart.EmptyCartResp, err error) {
	// Finish your business logic.
	err = model.EmptyCart(mysql.DB, s.ctx, req.UserId)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50001, "empty cart error")
	}

	return &cart.EmptyCartResp{}, nil
}
