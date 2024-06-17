//go:build wireinject
// +build wireinject

package provider

import (
	"github.com/google/wire"
	"go-shift/cmd/app/controller"
	"go-shift/cmd/app/repository"
	"go-shift/cmd/app/service"
)

type Initialization struct {
	authCtrl controller.AuthController
	authSvc  service.AuthService
	userRepo repository.UserRepository
}

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
	panic(wire.Build(NewInitialization, ProviderSet))
}
