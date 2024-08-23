package config

import (
	"github.com/redis/go-redis/v9"
	"go-shift/cmd/app/controller"
	"go-shift/cmd/app/controller/impl"
	"go-shift/cmd/app/repository"
	"go-shift/cmd/app/service"
	authServicePkg "go-shift/cmd/app/service/auth"
	apiService "go-shift/cmd/app/service/auth/api_service"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type Initialization struct {
	mysql   *gorm.DB
	mongodb *mongo.Client
	redis   *redis.Client

	UserRepository      repository.UserRepository
	AuthTokenRepository repository.AuthTokenRepository

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
	authTokenRepository := repository.ProvideAuthTokenRepository(mysql)

	redisService := service.ProvideRedisService(connectToRedis)
	workspaceService := service.ProvideWorkspaceService(workspaceRepository)

	authService := authServicePkg.ProvideAuthService(userRepository)
	oauthApiService := apiService.ProvideOauthApiService()
	googleOauthService := authServicePkg.ProvideGoogleOauthService(redisService, oauthApiService, userRepository, authTokenRepository)

	authController := controller.ProvideAuthController(authService)
	googleOauthController := impl.ProvideGoogleOauthController(googleOauthService)
	workspaceController := controller.ProvideWorkspaceController(workspaceService)

	return &Initialization{
		mysql:   mysql,
		mongodb: mongodb,
		redis:   connectToRedis,

		UserRepository:      userRepository,
		AuthTokenRepository: authTokenRepository,

		RedisService:       redisService,
		AuthService:        authService,
		GoogleOauthService: googleOauthService,

		AuthController:        authController,
		GoogleOauthController: googleOauthController,
		WorkspaceController:   workspaceController,
	}
}
