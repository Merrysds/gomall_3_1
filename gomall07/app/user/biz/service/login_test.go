package service

import (
	"context"
	"reflect"
	"testing"

	user "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/user"
)

func TestLogin_Run(t *testing.T) {
	ctx := context.Background()
	s := NewLoginService(ctx)
	// init req and assert value

	req := &user.LoginReq{}
	resp, err := s.Run(req)
	t.Logf("err: %v", err)
	t.Logf("resp: %v", resp)

	// todo: edit your unit test

}

func TestLoginService_Run(t *testing.T) {
	type args struct {
		req *user.LoginReq
	}
	tests := []struct {
		name     string
		s        *LoginService
		args     args
		wantResp *user.LoginResp
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.s.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("LoginService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
