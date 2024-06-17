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
	authCtrl controller.AuthController
	authSvc  service.AuthService
	userRepo repository.UserRepository
}

var db = wire.NewSet(config.ConnectToMysql)

func NewInitialization(
	authCtrl controller.AuthController,
	authService service.AuthService,
	userRepo repository.UserRepository,
) *Initialization {
	return &Initialization{
		authCtrl: authCtrl,
		authSvc:  authService,
		userRepo: userRepo,
	}
}

func Wire() *Initialization {
	panic(wire.Build(NewInitialization, db, ProviderSet))
}
