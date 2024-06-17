//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	"go-shift/cmd/app/controller"
	"go-shift/cmd/app/repository"
	"go-shift/cmd/app/service"
	"go-shift/config"
)

type Initialization struct {
	AuthCtrl controller.AuthController
	AuthSvc  service.AuthService
	UserRepo repository.UserRepository
}

var db = wire.NewSet(config.ConnectToMysql)

func NewInitialization(
	authCtrl controller.AuthController,
	authService service.AuthService,
	userRepo repository.UserRepository,
) *Initialization {
	return &Initialization{
		AuthCtrl: authCtrl,
		AuthSvc:  authService,
		UserRepo: userRepo,
	}
}

func Wire() *Initialization {
	panic(wire.Build(NewInitialization, db, ProviderSet))
}
