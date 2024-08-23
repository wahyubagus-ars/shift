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

	UserAccountRepository repository.UserRepository
	UserProfileRepository repository.UserProfileRepository
	AuthTokenRepository   repository.AuthTokenRepository

	RedisService       service.RedisService
	AuthService        authServicePkg.OAuthService
	OauthApiService    apiService.OauthApiService
	GoogleOauthService authServicePkg.OAuthService
	UserProfileService service.UserProfileService

	AuthController        controller.AuthController
	GoogleOauthController controller.AuthController
	WorkspaceController   controller.WorkspaceController
	UserProfileController controller.UserProfileController
}

func Init() *Initialization {
	mysql := ConnectToMysql()
	mongodb := ConnectToMongoDb()
	connectToRedis := ConnectToRedis()

	userAccountRepository := repository.ProvideUserRepository(mysql, mongodb)
	userProfileRepository := repository.ProvideUserProfileRepository(mysql)
	workspaceRepository := repository.ProvideWorkspaceRepository(mysql)
	authTokenRepository := repository.ProvideAuthTokenRepository(mysql)

	redisService := service.ProvideRedisService(connectToRedis)
	workspaceService := service.ProvideWorkspaceService(workspaceRepository)

	authService := authServicePkg.ProvideAuthService(userAccountRepository)
	oauthApiService := apiService.ProvideOauthApiService()
	googleOauthService := authServicePkg.ProvideGoogleOauthService(redisService, oauthApiService, userAccountRepository,
		&repository.UserProfileRepositoryImpl{},
		authTokenRepository)
	userProfileService := service.ProvideUserProfileService(userProfileRepository)

	authController := controller.ProvideAuthController(authService)
	googleOauthController := impl.ProvideGoogleOauthController(googleOauthService)
	workspaceController := controller.ProvideWorkspaceController(workspaceService)
	userProfileController := controller.ProvideUserProfileController(userProfileService)

	return &Initialization{
		mysql:   mysql,
		mongodb: mongodb,
		redis:   connectToRedis,

		UserAccountRepository: userAccountRepository,
		UserProfileRepository: userProfileRepository,
		AuthTokenRepository:   authTokenRepository,

		RedisService:       redisService,
		AuthService:        authService,
		GoogleOauthService: googleOauthService,
		UserProfileService: userProfileService,

		AuthController:        authController,
		GoogleOauthController: googleOauthController,
		WorkspaceController:   workspaceController,
		UserProfileController: userProfileController,
	}
}
