package service

import (
	"context"
	"strconv"

	common "github.com/cloudwego/gomall07/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/gomall07/app/frontend/infra/rpc"
	frontendUtils "github.com/cloudwego/gomall07/app/frontend/utils"
	"github.com/cloudwego/gomall07/rpc_gen/kitex_gen/cart"
	rpcproduct "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type GetCartService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewGetCartService(Context context.Context, RequestContext *app.RequestContext) *GetCartService {
	return &GetCartService{RequestContext: RequestContext, Context: Context}
}

func (h *GetCartService) Run(req *common.Empty) (resp map[string]any, err error) {
	var items []map[string]string
	cartResp, err := rpc.CartClient.GetCart(h.Context, &cart.GetCartReq{
		UserId: uint32(frontendUtils.GetUserIdFromCtx(h.Context)),
	})
	if err != nil {
		return nil, err
	}
	var total float64
	for _, item := range cartResp.Items {
		productResp, err := rpc.ProductClient.GetProduct(h.Context,
			&rpcproduct.GetProductReq{Id: item.ProductId})
		if err != nil {
			continue
		}
		if productResp.Product == nil {
			continue
		}
		p := productResp.Product
		items = append(items, map[string]string{
			"Name":        p.Name,
			"Description": p.Description,
			"Picture":     p.Picture,
			"Price":       strconv.FormatFloat(float64(p.Price), 'f', 2, 64),
			"Qty":         strconv.Itoa(int(item.Quantity)),
		})
		total += float64(item.Quantity) * float64(p.Price)
	}

	return utils.H{
		"title": "Cart",
		"items": items,
		"total": strconv.FormatFloat(float64(total), 'f', 2, 64),
	}, nil
}
