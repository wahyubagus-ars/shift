package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/constant"
	"go-shift/cmd/app/repository"
	"go-shift/pkg"
	"sync"
)

var (
	timeTrackingServiceOnce sync.Once
	timeTrackingService     *TimeTrackingServiceImpl
)

type TimeTrackingService interface {
	GetTimeEntries(c *gin.Context)
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

	c.JSON(200, timeEntries)
}

func ProvideTimeTrackingService(trackingRepository repository.TimeTrackingRepository) *TimeTrackingServiceImpl {
	timeTrackingServiceOnce.Do(func() {
		timeTrackingService = &TimeTrackingServiceImpl{
			timeTrackingRepository: trackingRepository,
		}
	})

	return timeTrackingService
}
