package service

import (
	"context"
	"reflect"
	"testing"

	cart "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/cart"
)

func TestGetCart_Run(t *testing.T) {
	ctx := context.Background()
	s := NewGetCartService(ctx)
	// init req and assert value

	req := &cart.GetCartReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestGetCartService_Run(t *testing.T) {
	type args struct {
		req *cart.GetCartReq
	}
	tests := []struct {
		name     string
		s        *GetCartService
		args     args
		wantResp *cart.GetCartResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCartService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("GetCartService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
