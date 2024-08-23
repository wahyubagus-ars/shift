package service

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/constant"
	"go-shift/cmd/app/domain/dao"
	"go-shift/cmd/app/domain/dto"
	"go-shift/cmd/app/repository"
	"go-shift/cmd/app/util"
	"go-shift/pkg"
	"sync"
)

var (
	userProfileService     *UserProfileServiceImpl
	userProfileServiceOnce sync.Once
)

type UserProfileService interface {
	GetUserProfile(c *gin.Context)
}

type UserProfileServiceImpl struct {
	userProfileRepository repository.UserProfileRepository
}

func (svc *UserProfileServiceImpl) GetUserProfile(c *gin.Context) {
	defer pkg.PanicHandler(c)

	token := c.Request.Header["Authorization-Token"]
	tokenPayload, _ := util.GetTokenPayload(token[0])

	userProfile, err := svc.userProfileRepository.FindByUserAccountEmail(tokenPayload.Email)
	if err != nil && userProfile == nil {
		pkg.PanicException(constant.DataNotFound)
	} else if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	apiResponse := dto.ApiResponse[dao.UserProfile]{
		ResponseKey:     constant.Success.GetResponseStatus(),
		ResponseMessage: constant.Success.GetResponseMessage(),
		Data:            *userProfile,
	}

	c.JSON(200, apiResponse)
}

func ProvideUserProfileService(profileRepository repository.UserProfileRepository) *UserProfileServiceImpl {
	userProfileServiceOnce.Do(func() {
		userProfileService = &UserProfileServiceImpl{
			userProfileRepository: profileRepository,
		}
	})

	return userProfileService
}
