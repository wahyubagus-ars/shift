package config

import (
	"go-shift/cmd/app/controller"
	"go-shift/cmd/app/repository"
	"go-shift/cmd/app/service"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Initialization struct {
	mysql   *gorm.DB
	mongodb *mongo.Client

	UserRepository repository.UserRepository

	AuthService        service.AuthService
	GoogleOauthService service.AuthService

	AuthController        controller.AuthController
	GoogleOauthController controller.AuthController
}

func Init() *Initialization {
	mysql := ConnectToMysql()
	mongodb := ConnectToMongoDb()

	userRepository := repository.ProvideUserRepository(mysql, mongodb)

	authService := service.ProvideAuthService(userRepository)
	googleOauthService := service.ProvideGoogleOauthService(userRepository)

	authController := controller.ProvideAuthController(authService)
	googleOauthController := controller.ProvideGoogleOauthController(googleOauthService)

	return &Initialization{
		mysql:   mysql,
		mongodb: mongodb,

		UserRepository: userRepository,

		AuthService:        authService,
		GoogleOauthService: googleOauthService,

		AuthController:        authController,
		GoogleOauthController: googleOauthController,
	}
}
