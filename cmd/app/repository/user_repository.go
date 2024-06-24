package repository

import (
	"go-shift/cmd/app/domain/dao"
	"go.mongodb.org/mongo-driver/mongo"
	"gorm.io/gorm"
	"sync"
	"time"
)

var (
	userRepositoy     *UserRepositoryImpl
	userRepositoyOnce sync.Once
)

type UserRepository interface {
	FindUserById(id int) int
	FindUserByEmail(email string) (dao.UserAccount, error)
	SaveInitiateUser(email string, authenticationId int) (dao.UserAccount, error)
}

type UserRepositoryImpl struct {
	mysql   *gorm.DB
	mongodb *mongo.Database
}

func (ur *UserRepositoryImpl) FindUserById(id int) int {
	return 1
}

func (ur *UserRepositoryImpl) FindUserByEmail(email string) (dao.UserAccount, error) {
	var user dao.UserAccount
	var err = ur.mysql.Where("email = ?", email).First(&user).Error

	if err != nil {
		return dao.UserAccount{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) SaveInitiateUser(email string, authenticationId int) (dao.UserAccount, error) {
	user := dao.UserAccount{
		Email:            email,
		AuthenticationID: authenticationId,
		BaseModel: dao.BaseModel{
			CreatedAt: time.Now(),
			CreatedBy: 0,
		},
	}

	if err := ur.mysql.Save(&user).Error; err != nil {
		return dao.UserAccount{}, err
	}

	return user, nil
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
