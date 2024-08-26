package controller

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/service"
	"sync"
)

var (
	timeTrackingControllerOnce sync.Once
	timeTrackingController     *TimeTrackingControllerImpl
)

type TimeTrackingController interface {
	GetTimeEntries(c *gin.Context)
	SubmitTimeEntry(c *gin.Context)
}

type TimeTrackingControllerImpl struct {
	timeTrackingService service.TimeTrackingService
}

func (ctrl *TimeTrackingControllerImpl) GetTimeEntries(c *gin.Context) {
	ctrl.timeTrackingService.GetTimeEntries(c)
}

func (ctrl *TimeTrackingControllerImpl) SubmitTimeEntry(c *gin.Context) {
	ctrl.timeTrackingService.SubmitTimeEntry(c)
}

func ProvideTimeTrackingController(trackingService service.TimeTrackingService) *TimeTrackingControllerImpl {
	timeTrackingControllerOnce.Do(func() {
		timeTrackingController = &TimeTrackingControllerImpl{
			timeTrackingService: trackingService,
		}
	})

	return timeTrackingController
}
