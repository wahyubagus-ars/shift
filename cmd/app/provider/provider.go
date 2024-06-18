package provider

import (
	"github.com/google/wire"
	"go-shift/cmd/app/controller"
	"go-shift/cmd/app/repository"
	"go-shift/cmd/app/service"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"sync"
)

var (
	ac     *controller.AuthControllerImpl
	acOnce sync.Once

	as     *service.AuthServiceImpl
	asOnce sync.Once

	ur     *repository.UserRepositoryImpl
	urOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		provideAuthController,
		provideAuthService,
		provideUserRepository,

		wire.Bind(new(controller.AuthController), new(*controller.AuthControllerImpl)),
		wire.Bind(new(service.AuthService), new(*service.AuthServiceImpl)),
		wire.Bind(new(repository.UserRepository), new(*repository.UserRepositoryImpl)),
	)
)

func provideAuthController(as service.AuthService) *controller.AuthControllerImpl {
	acOnce.Do(func() {
		ac = &controller.AuthControllerImpl{
			AuthSvc: as,
		}
	})

	return ac
}

func provideAuthService(ur repository.UserRepository) *service.AuthServiceImpl {
	asOnce.Do(func() {
		as = &service.AuthServiceImpl{
			UserRepo: ur,
		}
	})

	return as
}

func provideUserRepository(mysql *gorm.DB, mongo *mongo.Client) *repository.UserRepositoryImpl {
	urOnce.Do(func() {
		ur = &repository.UserRepositoryImpl{
			Mysql:   mysql,
			MongoDB: mongo.Database("shift_local"),
		}
	})

	return ur
}
