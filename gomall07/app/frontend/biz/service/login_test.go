package service

import (
	"testing"

	auth "github.com/cloudwego/gomall07/app/frontend/hertz_gen/frontend/auth"
)

func TestLoginService_Run(t *testing.T) {
	type args struct {
		req *auth.LoginReq
	}
	tests := []struct {
		name         string
		h            *LoginService
		args         args
		wantRedirect string
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRedirect, err := tt.h.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LoginService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRedirect != tt.wantRedirect {
				t.Errorf("LoginService.Run() = %v, want %v", gotRedirect, tt.wantRedirect)
			}
		})
	}
}
