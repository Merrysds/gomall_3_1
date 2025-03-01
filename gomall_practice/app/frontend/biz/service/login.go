package service

import (
	"context"
	"fmt"
	"time"

	auth "github.com/cloudwego/gomall07/app/frontend/hertz_gen/frontend/auth"
	"github.com/cloudwego/gomall07/app/frontend/infra/rpc"
	"github.com/cloudwego/gomall07/rpc_gen/kitex_gen/user"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol"
	"github.com/golang-jwt/jwt"
	"github.com/hertz-contrib/sessions"
)

type LoginService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewLoginService(Context context.Context, RequestContext *app.RequestContext) *LoginService {
	return &LoginService{RequestContext: RequestContext, Context: Context}
}

func (h *LoginService) Run(req *auth.LoginReq) (redirect string, err error) {
	resp, err := rpc.UserClient.Login(h.Context, &user.LoginReq{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return "", err
	}
	// 生成 JWT 令牌
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": resp.UserId,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // 设置过期时间为 24 小时
	})

	// 使用密钥签名令牌
	tokenString, err := token.SignedString([]byte("35783dfsghtuyidjhsgf")) // 请替换为你自己的密钥
	if err != nil {
		return "", err
	}

	// 将生成的 JWT 令牌设置为 Cookie 或者直接返回给前端
	h.RequestContext.SetCookie("jwt", tokenString, 86400, "/", "", protocol.CookieSameSiteLaxMode, false, true)
	fmt.Println("Generated JWT Token:", tokenString)

	session := sessions.Default(h.RequestContext)
	session.Set("user_id", resp.UserId)
	fmt.Println("User Info:", session)
	err = session.Save()
	if err != nil {
		return "", err
	}

	redirect = "/"
	if req.Next != "" {
		redirect = req.Next
	}
	return
}
