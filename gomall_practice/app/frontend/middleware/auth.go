package middleware

import (
	"context"

	frontendUtils "github.com/cloudwego/gomall07/app/frontend/utils"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/hertz-contrib/sessions"
)

func GlobalAuth() app.HandlerFunc { // 获取用户相关信息
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		ctx = context.WithValue(ctx, frontendUtils.SessionUserId, s.Get("user_id"))
		c.Next(ctx)
	}
}

func Auth() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		s := sessions.Default(c)
		userId := s.Get("user_id")

		if userId == nil {
			c.Redirect(302, []byte("/sign-in?next="+c.FullPath()))

			c.Abort()
			return
		}
		c.Next(ctx)
	}
}
