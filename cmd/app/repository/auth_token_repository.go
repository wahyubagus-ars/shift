package repository

import (
	"go-shift/cmd/app/domain/dao"
	"gorm.io/gorm"
	"sync"
)

var (
	authTokenRepository     *AuthTokenRepositoryImpl
	authTokenRepositoryOnce sync.Once
)

type AuthTokenRepository interface {
	SaveUserAuth(authToken *dao.AuthToken) (*dao.AuthToken, error)
}

type AuthTokenRepositoryImpl struct {
	mysql *gorm.DB
}

func (at *AuthTokenRepositoryImpl) SaveUserAuth(authToken *dao.AuthToken) (*dao.AuthToken, error) {
	if err := at.mysql.Save(authToken).Error; err != nil {
		return nil, err
	}

	return authToken, nil
}

func ProvideAuthTokenRepository(db *gorm.DB) *AuthTokenRepositoryImpl {
	authTokenRepositoryOnce.Do(func() {
		authTokenRepository = &AuthTokenRepositoryImpl{
			mysql: db,
		}
	})

	return authTokenRepository
}
