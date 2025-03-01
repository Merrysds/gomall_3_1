package service

import (
	"reflect"
	"testing"

	order "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/order"
)

func TestListOrderService_Run(t *testing.T) {
	type args struct {
		req *order.ListOrderReq
	}
	tests := []struct {
		name     string
		s        *ListOrderService
		args     args
		wantResp *order.ListOrderResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("ListOrderService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("ListOrderService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
