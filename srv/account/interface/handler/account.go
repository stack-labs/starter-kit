package handler

import (
	"context"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/test"
	"github.com/micro/go-micro/errors"
	"github.com/micro/go-micro/util/log"

	"github.com/micro-in-cn/starter-kit/srv/account/usecase"
	account "github.com/micro-in-cn/starter-kit/srv/pb/account"
)

type Account struct {
	userUsecase usecase.UserUsecase
}

func NewAccountService(userUsecase usecase.UserUsecase) *Account {
	return &Account{
		userUsecase: userUsecase,
	}
}

// Call is a single request handler called via client.Call or the generated client code
func (a *Account) Login(ctx context.Context, req *account.LoginRequest, rsp *account.LoginResponse) error {
	log.Log("Received Example.Call request")

	user, err := a.userUsecase.LoginUser(req.Username, req.Password)
	if err != nil {
		return err
	} else if user == nil {
		return errors.New("go.micro.srv.account", "用户名或密码错误", 200)
	}

	claims := jwt.StandardClaims{
		Id:        req.Username,
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
	}

	privateKey := test.LoadRSAPrivateKeyFromDisk("./conf/auth_key")
	tokenString := test.MakeSampleToken(claims, privateKey)

	rsp.Token = tokenString

	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (*Account) Logout(ctx context.Context, req *account.Request, rsp *account.LogoutResponse) error {
	log.Log("Received Example.Call request")
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (*Account) Info(ctx context.Context, req *account.Request, rsp *account.InfoResponse) error {
	log.Log("Received Example.Call request")
	rsp.Name = "Hobo"
	rsp.Avatar = "https://avatars3.githubusercontent.com/u/730866?s=460&v=4"
	return nil
}
