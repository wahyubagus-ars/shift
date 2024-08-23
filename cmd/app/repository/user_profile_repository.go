package repository

import (
	"go-shift/cmd/app/domain/dao/table"
	"gorm.io/gorm"
	"sync"
)

var (
	userProfileRepository     *UserProfileRepositoryImpl
	userProfileRepositoryOnce sync.Once
)

type UserProfileRepository interface {
	FindByUserAccountId(id int) (*table.UserProfile, error)
	FindByUserAccountEmail(email string) (*table.UserProfile, error)
	SaveUserProfile(profile *table.UserProfile) (*table.UserProfile, error)
}

type UserProfileRepositoryImpl struct {
	mysql *gorm.DB
}

func (r *UserProfileRepositoryImpl) FindByUserAccountId(id int) (*table.UserProfile, error) {
	var userProfile table.UserProfile
	err := r.mysql.Where("user_account_id", id).Find(&userProfile).Error

	if err != nil {
		return nil, err
	}

	return &userProfile, nil
}

func (r *UserProfileRepositoryImpl) FindByUserAccountEmail(email string) (*table.UserProfile, error) {
	var userProfile table.UserProfile
	err := r.mysql.Raw("SELECT up.* FROM user_profile as up "+
		"JOIN user_account as ua ON up.user_account_id = ua.id "+
		"WHERE ua.email = ?", email).Find(&userProfile).Error

	if err != nil {
		return nil, err
	}

	return &userProfile, nil
}

func (r *UserProfileRepositoryImpl) SaveUserProfile(profile *table.UserProfile) (*table.UserProfile, error) {
	return nil, nil
}

func ProvideUserProfileRepository(mysql *gorm.DB) *UserProfileRepositoryImpl {
	userProfileRepositoryOnce.Do(func() {
		userProfileRepository = &UserProfileRepositoryImpl{
			mysql: mysql,
		}
	})

	return userProfileRepository
}
