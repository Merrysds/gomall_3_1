package service

import (
	"context"
	"reflect"
	"testing"

	order "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/order"
)

func TestPlaceOrder_Run(t *testing.T) {
	ctx := context.Background()
	s := NewPlaceOrderService(ctx)
	// init req and assert value

	req := &order.PlaceOrderReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestPlaceOrderService_Run(t *testing.T) {
	type args struct {
		req *order.PlaceOrderReq
	}
	tests := []struct {
		name     string
		s        *PlaceOrderService
		args     args
		wantResp *order.PlaceOrderResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("PlaceOrderService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("PlaceOrderService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
