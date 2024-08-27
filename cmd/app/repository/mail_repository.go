package repository

import (
	"context"
	"go-shift/cmd/app/domain/dao/collection"
	"go-shift/cmd/app/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var (
	mailRepositoryOnce sync.Once
	mailRepository     *MailRepositoryImpl
)

type MailRepository interface {
	SaveVerificationEmail(invitation *collection.EmailVerification) (*collection.EmailVerification, error)
	FindVerificationEmailById(id string) (*collection.EmailVerification, error)
}

type MailRepositoryImpl struct {
	mongodb *mongo.Database
}

func (r *MailRepositoryImpl) SaveVerificationEmail(invitation *collection.EmailVerification) (*collection.EmailVerification, error) {
	result, err := r.mongodb.Collection("verificationEmail").InsertOne(context.TODO(), &invitation)

	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		invitation.ID = oid
	}

	return invitation, nil
}

func (r *MailRepositoryImpl) FindVerificationEmailById(id string) (*collection.EmailVerification, error) {
	var emailVerification collection.EmailVerification
	objectId, err := util.GenerateObjectId(id)
	err = r.mongodb.Collection("verificationEmail").
		FindOne(context.TODO(), bson.M{"_id": objectId}).
		Decode(&emailVerification)
	if err != nil {
		return nil, err
	}

	return &emailVerification, nil
}

func ProvideMailRepository(mongodb *mongo.Client) *MailRepositoryImpl {
	mailRepositoryOnce.Do(func() {
		mailRepository = &MailRepositoryImpl{
			mongodb: mongodb.Database("shiftLocal"),
		}
	})

	return mailRepository
}
