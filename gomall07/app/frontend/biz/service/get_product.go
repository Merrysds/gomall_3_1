package service

import (
	"context"

	product "github.com/cloudwego/gomall07/app/frontend/hertz_gen/frontend/product"
	"github.com/cloudwego/gomall07/app/frontend/infra/rpc"
	rpcproduct "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetProductService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetProductService(Context context.Context, RequestContext *app.RequestContext) *GetProductService {
	return &GetProductService{RequestContext: RequestContext, Context: Context}
}

func (h *GetProductService) Run(req *product.ProductReq) (resp map[string]any, err error) {
	p, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: req.GetId()})
	if err != nil {
		return nil, err
	}
	return utils.H{
		"item": p.Product,
	}, nil
}
