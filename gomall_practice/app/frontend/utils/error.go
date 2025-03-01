package utils

import "github.com/cloudwego/hertz/pkg/common/hlog"

func MustHandleErr(err error) {
	if err != nil {
		hlog.Fatal(err)
	}
}
