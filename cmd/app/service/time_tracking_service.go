package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/constant"
	"go-shift/cmd/app/domain/dao/collection"
	"go-shift/cmd/app/domain/dto"
	"go-shift/cmd/app/domain/dto/system"
	"go-shift/cmd/app/repository"
	"go-shift/cmd/app/util"
	"go-shift/pkg"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"sync"
)

var (
	timeTrackingServiceOnce sync.Once
	timeTrackingService     *TimeTrackingServiceImpl
)

type TimeTrackingService interface {
	GetTimeEntries(c *gin.Context)
	SubmitTimeEntry(c *gin.Context)
}

type TimeTrackingServiceImpl struct {
	timeTrackingRepository repository.TimeTrackingRepository
}

func (svc *TimeTrackingServiceImpl) GetTimeEntries(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute get time entries")

	timeEntries, err := svc.timeTrackingRepository.GetTimeEntries(1, 1)

	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	apiResponse := &system.ApiResponse[[]collection.TimeEntry]{
		ResponseKey:     constant.Success.GetResponseStatus(),
		ResponseMessage: constant.Success.GetResponseMessage(),
		Data:            timeEntries,
	}
	c.JSON(200, apiResponse)
}

func (svc *TimeTrackingServiceImpl) SubmitTimeEntry(c *gin.Context) {
	defer pkg.PanicHandler(c)

	timeEntryReq := dto.TimeEntryRequest{}

	if err := c.ShouldBindJSON(&timeEntryReq); err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	timeEntryCollection := &collection.TimeEntry{
		UserID:      1,
		ProjectID:   timeEntryReq.ProjectId,
		TaskID:      primitive.NewObjectID(),
		StartTime:   timeEntryReq.StartTime,
		EndTime:     timeEntryReq.EndTime,
		Duration:    timeEntryReq.Duration, // Duration in minutes (8 hours)
		Title:       timeEntryReq.Title,
		Description: timeEntryReq.Description,
		IsBillable:  timeEntryReq.IsBillable,
		BaseModel: collection.BaseModel{
			CreatedAt: util.GenerateTimePtr(),
			CreatedBy: util.GenerateIntPtr(1),
		},
	}

	timeEntry, err := svc.timeTrackingRepository.SubmitTimeEntry(timeEntryCollection)

	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(200, timeEntry)

}

func ProvideTimeTrackingService(trackingRepository repository.TimeTrackingRepository) *TimeTrackingServiceImpl {
	timeTrackingServiceOnce.Do(func() {
		timeTrackingService = &TimeTrackingServiceImpl{
			timeTrackingRepository: trackingRepository,
		}
	})

	return timeTrackingService
}
