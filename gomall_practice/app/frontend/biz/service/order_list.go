package service

import (
	"context"
	"time"

	common "github.com/cloudwego/gomall07/app/frontend/hertz_gen/frontend/common"
	"github.com/cloudwego/gomall07/app/frontend/infra/rpc"
	"github.com/cloudwego/gomall07/app/frontend/types"
	frontendUtils "github.com/cloudwego/gomall07/app/frontend/utils"
	"github.com/cloudwego/gomall07/rpc_gen/kitex_gen/order"
	rpcproduct "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/product"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
)

type OrderListService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewOrderListService(Context context.Context, RequestContext *app.RequestContext) *OrderListService {
	return &OrderListService{RequestContext: RequestContext, Context: Context}
}

func (h *OrderListService) Run(req *common.Empty) (resp map[string]any, err error) {
	//defer func() {
	// hlog.CtxInfof(h.Context, "req = %+v", req)
	// hlog.CtxInfof(h.Context, "resp = %+v", resp)
	//}()
	// todo edit your code
	userId := frontendUtils.GetUserIdFromCtx(h.Context)
	orderResp, err := rpc.OrderClient.ListOrder(h.Context, &order.ListOrderReq{UserId: uint32(userId)})
	if err != nil {
		return nil, err
	}

	var list []types.Order
	for _, v := range orderResp.Orders {
		var total float32
		var items []types.OrderItem

		for _, v := range v.Items {
			total += v.Cost
			i := v.Item
			productResp, err := rpc.ProductClient.GetProduct(h.Context, &rpcproduct.GetProductReq{Id: i.ProductId}) // 这句话打代码的时候没有响应
			if err != nil {
				return nil, err
			}
			if productResp == nil || productResp.Product == nil {
				continue
			}
			p := productResp.Product

			items = append(items, types.OrderItem{
				ProductName: p.Name,
				Picture:     p.Picture,
				Cost:        v.Cost,
				Qty:         i.Quantity,
			})
		}

		created := time.Unix(int64(v.CreatedAt), 0)
		list = append(list, types.Order{
			OrderId:     v.OrderId,
			CreatedDate: created.Format("2025-02-25 10:34"),
			Cost:        total,
			Items:       items,
		})
	}
	return utils.H{
		"title":  "Order",
		"orders": list,
	}, nil
}
