package config

import (
	"github.com/redis/go-redis/v9"
	"go-shift/cmd/app/controller"
	"go-shift/cmd/app/controller/impl"
	"go-shift/cmd/app/repository"
	"go-shift/cmd/app/service"
	apiService "go-shift/cmd/app/service/api_service"
	authServicePkg "go-shift/cmd/app/service/auth"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Initialization struct {
	mysql   *gorm.DB
	mongodb *mongo.Client
	redis   *redis.Client

	UserRepository repository.UserRepository

	RedisService       service.RedisService
	AuthService        authServicePkg.AuthService
	OauthApiService    apiService.OauthApiService
	GoogleOauthService authServicePkg.AuthService

	AuthController        controller.AuthController
	GoogleOauthController controller.AuthController
	WorkspaceController   controller.WorkspaceController
}

func Init() *Initialization {
	mysql := ConnectToMysql()
	mongodb := ConnectToMongoDb()
	connectToRedis := ConnectToRedis()

	userRepository := repository.ProvideUserRepository(mysql, mongodb)
	workspaceRepository := repository.ProvideWorkspaceRepository(mysql)

	redisService := service.ProvideRedisService(connectToRedis)
	workspaceService := service.ProvideWorkspaceService(workspaceRepository)

	authService := authServicePkg.ProvideAuthService(userRepository)
	oauthApiService := apiService.ProvideOauthApiService()
	googleOauthService := authServicePkg.ProvideGoogleOauthService(redisService, oauthApiService, userRepository)

	authController := controller.ProvideAuthController(authService)
	googleOauthController := impl.ProvideGoogleOauthController(googleOauthService)
	workspaceController := controller.ProvideWorkspaceController(workspaceService)

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
		WorkspaceController:   workspaceController,
	}
}
