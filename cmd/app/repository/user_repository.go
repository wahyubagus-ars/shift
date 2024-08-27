package repository

import (
	"go-shift/cmd/app/domain/dao/table"
	"go-shift/cmd/app/util"
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
	FindUserByEmail(email string) (table.UserAccount, error)
	SaveInitiateUser(email string, authenticationId int) (table.UserAccount, error)
	UpdateEmailVerifiedAt(email string) error
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
			CreatedBy: util.GenerateIntPtr(0),
		},
	}

	if err := ur.mysql.Save(&user).Error; err != nil {
		return table.UserAccount{}, err
	}

	return user, nil
}

func (ur *UserRepositoryImpl) UpdateEmailVerifiedAt(email string) error {
	err := ur.mysql.Model(&table.UserAccount{}).Where("email = ?", email).
		Update("email_verified_at", time.Now()).Error

	if err != nil {
		return err
	}

	return nil
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
