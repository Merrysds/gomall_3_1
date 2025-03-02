package service

import (
	"reflect"
	"testing"

	checkout "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/checkout"
)

func TestCheckoutService_Run(t *testing.T) {
	type args struct {
		req *checkout.CheckoutReq
	}
	tests := []struct {
		name     string
		s        *CheckoutService
		args     args
		wantResp *checkout.CheckoutResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("CheckoutService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("CheckoutService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
