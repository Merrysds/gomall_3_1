package service

import (
	"context"
	"reflect"
	"testing"

	payment "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/payment"
)

func TestCharge_Run(t *testing.T) {
	ctx := context.Background()
	s := NewChargeService(ctx)
	// init req and assert value

	req := &payment.ChargeReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestChargeService_Run(t *testing.T) {
	type args struct {
		req *payment.ChargeReq
	}
	tests := []struct {
		name     string
		s        *ChargeService
		args     args
		wantResp *payment.ChargeResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChargeService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ChargeService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
