package middleware

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/golang-jwt/jwt"
)

func Register(h *server.Hertz) {
	h.Use(GlobalAuth())
}

func JWT() app.HandlerFunc {
	return func(c context.Context, ctx *app.RequestContext) {
		if ctx.FullPath() == "" {
			return
		}
		whitelist := []string{
			"/sign-in", "/sign-up", "/", "/auth/register", "/auth/login", "/auth/logout",
		}
		isWhitelisted := false
		for _, path := range whitelist {
			if ctx.FullPath() == path {
				isWhitelisted = true
				break
			}
		}

		// 对于 /about 路径，允许继续处理并尝试解析 JWT
		if ctx.FullPath() == "/about" || strings.HasPrefix(ctx.FullPath(), "/category") ||
			strings.HasPrefix(ctx.FullPath(), "/product") || strings.HasPrefix(ctx.FullPath(), "/search") {

			fmt.Println("进入特殊路径：", ctx.FullPath())
			tokenString := ctx.Cookie("jwt")
			if string(tokenString) != "" {
				fmt.Println("特殊路径token不为空")
				token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
					return []byte(os.Getenv("JWT_SECRET")), nil
				})
				if err == nil && token.Valid {
					claims := token.Claims.(jwt.MapClaims)
					userID := claims["user_id"]
					ctx.Set("user_id", userID)
					fmt.Println("middle_ware_user_id", userID)
				}
			}
			ctx.Next(c)
			ctx.Abort()
			return
		}

		// 对于其他白名单路径，直接放行
		if isWhitelisted {
			fmt.Println("进入白名单+路径：", ctx.FullPath())

			fmt.Println("white!")
			ctx.Next(c)
			ctx.Abort()
			return
		}

		// 对于非白名单路径，验证 JWT
		tokenString := ctx.Cookie("jwt")
		if string(tokenString) == "" {
			fmt.Println("没有token")
			fmt.Println("没有tokenpath：", ctx.FullPath())
			ctx.HTML(401, "error", map[string]interface{}{
				"ErrorMessage": "请先注册登录",
			})
			ctx.Abort()
			return
		}

		token, err := jwt.Parse(string(tokenString), func(token *jwt.Token) (interface{}, error) {
			fmt.Println("token:", string(tokenString))
			fmt.Println("密钥：", os.Getenv("JWT_SECRET"))
			return []byte(os.Getenv("JWT_SECRET")), nil
		})

		if err != nil || !token.Valid {

			// 重定向到登录页面，并显示提示信息
			fmt.Println("token验证不对")
			fmt.Println("token验证不对path", ctx.FullPath())
			ctx.HTML(401, "error", map[string]interface{}{
				"ErrorMessage": "token校验失败,可能存在登录身份过期,请重新登录",
			})
			ctx.Abort()
			return

		}

		claims := token.Claims.(jwt.MapClaims)
		userID := claims["user_id"]
		ctx.Set("user_id", userID)
		fmt.Println("middle_ware_user_id", userID)

		ctx.Next(c)
		ctx.Abort()
	}
}
