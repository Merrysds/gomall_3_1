package service

import (
	"context"

	"github.com/cloudwego/gomall07/app/cart/biz/dal/mysql"
	"github.com/cloudwego/gomall07/app/cart/biz/model"
	"github.com/cloudwego/gomall07/app/cart/rpc"
	cart "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/cart"
	"github.com/cloudwego/gomall07/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/kitex/pkg/kerrors"
)

type AddItemService struct {
	ctx context.Context
} // NewAddItemService new AddItemService
func NewAddItemService(ctx context.Context) *AddItemService {
	return &AddItemService{ctx: ctx}
}

// Run create note info
func (s *AddItemService) Run(req *cart.AddItemReq) (resp *cart.AddItemResp, err error) {
	// Finish your business logic.
	// Finish your business logic.
	getProduct, err := rpc.ProductClient.GetProduct(s.ctx, &product.GetProductReq{Id: req.Item.ProductId})
	if err != nil {
		return nil, err
	}

	if getProduct.Product == nil || getProduct.Product.Id == 0 {
		return nil, kerrors.NewBizStatusError(40004, "product not exist")
	}

	cartItem := &model.Cart{
		UserId:    req.UserId,
		ProductId: req.Item.ProductId,
		Qty:       req.Item.Quantity,
	}

	err = model.AddItem(s.ctx, mysql.DB, cartItem)
	if err != nil {
		return nil, kerrors.NewBizStatusError(50000, err.Error())
	}
	return &cart.AddItemResp{}, nil
}
