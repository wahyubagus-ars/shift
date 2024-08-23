package repository

import (
	"go-shift/cmd/app/domain/dao/table"
	"go-shift/cmd/app/util"
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
	FindUserByEmail(email string) (table.UserAccount, error)
	SaveInitiateUser(email string, authenticationId int) (table.UserAccount, error)
}

type UserRepositoryImpl struct {
	mysql   *gorm.DB
	mongodb *mongo.Database
}

func (ur *UserRepositoryImpl) FindUserById(id int) int {
	return 1
}

func (ur *UserRepositoryImpl) FindUserByEmail(email string) (table.UserAccount, error) {
	var user table.UserAccount
	var err = ur.mysql.Where("email = ?", email).First(&user).Error

	if err != nil {
		return table.UserAccount{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) SaveInitiateUser(email string, authenticationId int) (table.UserAccount, error) {

	user := table.UserAccount{
		Email:            email,
		AuthenticationID: authenticationId,
		BaseModel: table.BaseModel{
			CreatedAt: util.GenerateTimePtr(),
			CreatedBy: util.IntPtr(0),
		},
	}

	if err := ur.mysql.Save(&user).Error; err != nil {
		return table.UserAccount{}, err
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
