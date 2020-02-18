package handler

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/test"
	"github.com/micro-in-cn/starter-kit/console/account/conf"
	"github.com/micro/go-micro/v2/errors"
	"github.com/micro/go-micro/v2/util/log"

	account "github.com/micro-in-cn/starter-kit/console/account/genproto/srv"
	"github.com/micro-in-cn/starter-kit/console/account/usecase"
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
	log.Log("Received Account.Login request")

	user, err := a.userUsecase.LoginUser(req.Username, req.Password)
	if err != nil {
		return err
	} else if user == nil {
		return errors.New("go.micro.srv.account", "用户名或密码错误", 200)
	}

	claims := jwt.StandardClaims{
		Id:        strconv.FormatInt(user.Id, 10),
		ExpiresAt: time.Now().Add(time.Minute * 10).Unix(),
		Subject:   req.Username,
	}

	privateKey := test.LoadRSAPrivateKeyFromDisk(conf.BASE_PATH + "auth_key")
	tokenString := test.MakeSampleToken(claims, privateKey)

	rsp.Token = tokenString

	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (*Account) Logout(ctx context.Context, req *account.Request, rsp *account.LogoutResponse) error {
	log.Log("Received Account.Logout request")
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (a *Account) Info(ctx context.Context, req *account.Request, rsp *account.InfoResponse) error {
	log.Log("Received Account.Info request")
	user, err := a.userUsecase.GetUser(req.Id)
	if err != nil {
		return err
	}

	rsp.Name = fmt.Sprintf("%s-ID:%d", user.Name, req.Id)
	rsp.Avatar = "https://avatars3.githubusercontent.com/u/730866?s=460&v=4"
	return nil
}
