package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"sync"
)

var (
	userRepositoy     *UserRepositoryImpl
	userRepositoyOnce sync.Once
)

type UserRepository interface {
	FindUserById(id int) int
}

type UserRepositoryImpl struct {
	mysql   *gorm.DB
	mongodb *mongo.Database
}

func (ur *UserRepositoryImpl) FindUserById(id int) int {
	return 1
}

func ProvideUserRepository(mysql *gorm.DB, mongo *mongo.Client) *UserRepositoryImpl {
	userRepositoyOnce.Do(func() {
		userRepositoy = &UserRepositoryImpl{
			mysql:   mysql,
			mongodb: mongo.Database("shift_local"),
		}
	})

	return userRepositoy
}
