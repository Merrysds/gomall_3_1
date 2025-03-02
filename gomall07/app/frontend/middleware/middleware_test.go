package middleware

import (
	"reflect"
	"testing"

	"github.com/cloudwego/hertz/pkg/app"
)

func TestJWT(t *testing.T) {
	tests := []struct {
		name string
		want app.HandlerFunc
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := JWT(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("JWT() = %v, want %v", got, tt.want)
			}
		})
	}
}
