package service

import (
	"context"

	category "github.com/cloudwego/gomall07/app/frontend/hertz_gen/frontend/category"
	"github.com/cloudwego/gomall07/app/frontend/infra/rpc"
	rpcproduct "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type CategoryService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewCategoryService(Context context.Context, RequestContext *app.RequestContext) *CategoryService {
	return &CategoryService{RequestContext: RequestContext, Context: Context}
}

func (h *CategoryService) Run(req *category.CategoryReq) (resp map[string]any, err error) {
	p, _ := rpc.ProductClient.ListProducts(h.Context, &rpcproduct.ListProductsReq{CategoryName: req.Category})
	return utils.H{
		"title": "Category",
		"items": p.Products,
	}, nil
}
