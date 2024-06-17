package repository

import "gorm.io/gorm"

type UserRepository interface {
	FindUserById(id int) int
}

type UserRepositoryImpl struct {
	Db *gorm.DB
}

func (ur *UserRepositoryImpl) FindUserById(id int) int {
	return 1
}
