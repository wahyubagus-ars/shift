package repository

type UserRepository interface {
	FindUserById(id int) int
}

type UserRepositoryImpl struct {
}

func (ur *UserRepositoryImpl) FindUserById(id int) int {
	return 1
}
