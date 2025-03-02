package service

import (
	"reflect"
	"testing"

	common "github.com/cloudwego/gomall07/app/frontend/hertz_gen/frontend/common"
)

func TestLogoutService_Run(t *testing.T) {
	type args struct {
		req *common.Empty
	}
	tests := []struct {
		name     string
		h        *LogoutService
		args     args
		wantResp *common.Empty
		wantErr  bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotResp, err := tt.h.Run(tt.args.req)
			if (err != nil) != tt.wantErr {
				t.Errorf("LogoutService.Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotResp, tt.wantResp) {
				t.Errorf("LogoutService.Run() = %v, want %v", gotResp, tt.wantResp)
			}
		})
	}
}
