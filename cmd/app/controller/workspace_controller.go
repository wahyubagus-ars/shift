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
}

type WorkspaceControllerImpl struct {
	workspaceService service.WorkspaceService
}

func (wc *WorkspaceControllerImpl) GetWorkspace(c *gin.Context) {
	wc.workspaceService.GetWorkspace()
}

func ProvideWorkspaceController(workspaceService service.WorkspaceService) *WorkspaceControllerImpl {
	workspaceControllerOnce.Do(func() {
		workspaceController = &WorkspaceControllerImpl{
			workspaceService: workspaceService,
		}
	})

	return workspaceController
}
