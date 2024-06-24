package config

import (
	"github.com/redis/go-redis/v9"
	"go-shift/cmd/app/controller"
	"go-shift/cmd/app/repository"
	"go-shift/cmd/app/service"
	apiservice "go-shift/cmd/app/service/api_service"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Initialization struct {
	mysql   *gorm.DB
	mongodb *mongo.Client
	redis   *redis.Client

	UserRepository repository.UserRepository

	RedisService       service.RedisService
	AuthService        service.AuthService
	OauthApiService    apiservice.OauthApiService
	GoogleOauthService service.AuthService

	AuthController        controller.AuthController
	GoogleOauthController controller.AuthController
}

func Init() *Initialization {
	mysql := ConnectToMysql()
	mongodb := ConnectToMongoDb()
	connectToRedis := ConnectToRedis()

	userRepository := repository.ProvideUserRepository(mysql, mongodb)

	redisService := service.ProvideRedisService(connectToRedis)
	authService := service.ProvideAuthService(userRepository)
	oauthApiService := apiservice.ProvideOauthApiService()
	googleOauthService := service.ProvideGoogleOauthService(redisService, oauthApiService, userRepository)

	authController := controller.ProvideAuthController(authService)
	googleOauthController := controller.ProvideGoogleOauthController(googleOauthService)

	return &Initialization{
		mysql:   mysql,
		mongodb: mongodb,
		redis:   connectToRedis,

		UserRepository: userRepository,

		RedisService:       redisService,
		AuthService:        authService,
		GoogleOauthService: googleOauthService,

		AuthController:        authController,
		GoogleOauthController: googleOauthController,
	}
}
