package repository

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/domain/dao/collection"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"sync"
)

var (
	timeTrackingRepositoryOnce sync.Once
	timeTrackingRepository     *TimeTrackingRepositoryImpl
)

type TimeTrackingRepository interface {
	SubmitTimeEntry(timeEntry *collection.TimeEntry) (*collection.TimeEntry, error)
	GetTimeEntries(userId int, projectId int) ([]collection.TimeEntry, error)
}

type TimeTrackingRepositoryImpl struct {
	mongodb *mongo.Database
}

func (r *TimeTrackingRepositoryImpl) GetTimeEntries(userId int, projectId int) ([]collection.TimeEntry, error) {
	timeEntriesDocs, err := r.mongodb.Collection("timeEntry").Find(context.TODO(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}
	defer func(timeEntriesDocs *mongo.Cursor, ctx context.Context) {
		err := timeEntriesDocs.Close(ctx)
		if err != nil {

		}
	}(timeEntriesDocs, context.TODO())

	var timeEntries []collection.TimeEntry
	for timeEntriesDocs.Next(context.TODO()) {
		var timeEntry collection.TimeEntry
		if err = timeEntriesDocs.Decode(&timeEntry); err != nil {
			log.Fatal(err)
		}
		timeEntries = append(timeEntries, timeEntry)
	}

	return timeEntries, nil
}

func (r *TimeTrackingRepositoryImpl) SubmitTimeEntry(timeEntry *collection.TimeEntry) (*collection.TimeEntry, error) {
	result, err := r.mongodb.Collection("timeEntry").InsertOne(context.TODO(), &timeEntry)

	if err != nil {
		return nil, err
	}

	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		timeEntry.ID = oid
	}

	return timeEntry, nil
}

func ProvideTimeTrackingRepository(mongodb *mongo.Client) *TimeTrackingRepositoryImpl {
	timeTrackingRepositoryOnce.Do(func() {
		timeTrackingRepository = &TimeTrackingRepositoryImpl{
			mongodb: mongodb.Database("shiftLocal"),
		}
	})

	return timeTrackingRepository
}
