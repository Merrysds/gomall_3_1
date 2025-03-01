package main

import (
	"context"
	payment "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/payment"
	"github.com/cloudwego/gomall07/app/payment/biz/service"
)

// PaymentServiceImpl implements the last service interface defined in the IDL.
type PaymentServiceImpl struct{}

// Charge implements the PaymentServiceImpl interface.
func (s *PaymentServiceImpl) Charge(ctx context.Context, req *payment.ChargeReq) (resp *payment.ChargeResp, err error) {
	resp, err = service.NewChargeService(ctx).Run(req)

	return resp, err
}
