package controller

import (
	"github.com/gin-gonic/gin"
	"go-shift/cmd/app/service"
	"sync"
)

var (
	workspaceController     *WorkspaceControllerImpl
	workspaceControllerOnce sync.Once
)

type WorkspaceController interface {
	GetWorkspace(c *gin.Context)
	CreateWorkspace(c *gin.Context)
}

type WorkspaceControllerImpl struct {
	workspaceService service.WorkspaceService
}

func (controller *WorkspaceControllerImpl) GetWorkspace(c *gin.Context) {
	controller.workspaceService.GetWorkspace()
}

func (controller *WorkspaceControllerImpl) CreateWorkspace(c *gin.Context) {
	controller.workspaceService.CreateWorkspace(c)
}

func ProvideWorkspaceController(workspaceService service.WorkspaceService) *WorkspaceControllerImpl {
	workspaceControllerOnce.Do(func() {
		workspaceController = &WorkspaceControllerImpl{
			workspaceService: workspaceService,
		}
	})

	return workspaceController
}
