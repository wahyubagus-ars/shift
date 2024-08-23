package controller

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/service"
	"sync"
)

var (
	userProfileControllerOnce sync.Once
	userProfileController     *UserProfileControllerImpl
)

type UserProfileController interface {
	GetUserProfile(c *gin.Context)
}

type UserProfileControllerImpl struct {
	userProfileService service.UserProfileService
}

func (ctrl *UserProfileControllerImpl) GetUserProfile(c *gin.Context) {
	ctrl.userProfileService.GetUserProfile(c)
}

func ProvideUserProfileController(profileService service.UserProfileService) *UserProfileControllerImpl {
	userProfileControllerOnce.Do(func() {
		userProfileController = &UserProfileControllerImpl{
			userProfileService: profileService,
		}
	})

	return userProfileController
}
