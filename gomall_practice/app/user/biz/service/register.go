package service

import (
	"context"
	"errors"

	"github.com/cloudwego/gomall07/app/user/biz/dal/mysql"
	"github.com/cloudwego/gomall07/app/user/biz/model"
	user "github.com/cloudwego/gomall07/rpc_gen/kitex_gen/user"
	"golang.org/x/crypto/bcrypt"
)

type RegisterService struct {
	ctx context.Context
} // NewRegisterService new RegisterService
func NewRegisterService(ctx context.Context) *RegisterService {
	return &RegisterService{ctx: ctx}
}

// Run create note info
func (s *RegisterService) Run(req *user.RegisterReq) (resp *user.RegisterResp, err error) {
	// 检查请求数据
	if req.Email == "" || req.Password == "" || req.PasswordConfirm == "" {
		return nil, errors.New("email or password is empty")
	}

	if req.Password != req.PasswordConfirm {
		return nil, errors.New("password not match")
	}

	// 密码加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	// 创建用户
	newUser := &model.User{
		Email:          req.Email,
		PasswordHashed: string(hashedPassword),
	}

	// 保存用户到数据库
	if err = model.Create(mysql.DB, newUser); err != nil {
		return nil, err
	}

	// 返回用户信息（ID），方便前端生成 JWT
	return &user.RegisterResp{UserId: int32(newUser.ID)}, nil
}
