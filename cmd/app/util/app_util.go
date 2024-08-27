package util

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func GenerateIntPtr(i int) *int {
	return &i
}

func GenerateTimePtr() *time.Time {
	currentTime := time.Now()

	return &currentTime
}

func GenerateObjectId(id string) (primitive.ObjectID, error) {
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return objectID, err
	}

	return objectID, nil
}
