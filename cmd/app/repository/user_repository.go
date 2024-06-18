package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindUserById(id int) int
}

type UserRepositoryImpl struct {
	Mysql   *gorm.DB
	MongoDB *mongo.Database
}

func (ur *UserRepositoryImpl) FindUserById(id int) int {
	return 1
}
