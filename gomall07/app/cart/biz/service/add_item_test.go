package service

import (
	"reflect"
	"testing"

	cart "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/cart"
)

func TestAddItemService_Run(t *testing.T) {
	type args struct {
		req *cart.AddItemReq
	}
	tests := []struct {
		name     string
		s        *AddItemService
		args     args
		wantResp *cart.AddItemResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddItemService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("AddItemService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
