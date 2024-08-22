package service

import (
	"go-shift/cmd/app/repository"
	"sync"
)

var (
	workspaceService     *WorkspaceServiceImpl
	workspaceServiceOnce sync.Once
)

type WorkspaceService interface {
	GetWorkspace()
	AddWorkspace()
	//UpdateWorkspace()
	//DeleteWorkspace()
}

type WorkspaceServiceImpl struct {
	workspaceRepository repository.WorkspaceRepository
}

func (ws *WorkspaceServiceImpl) GetWorkspace() {
	ws.workspaceRepository.GetWorkspace()
}

func (ws *WorkspaceServiceImpl) AddWorkspace() {

}

func ProvideWorkspaceService(workspaceRepository repository.WorkspaceRepository) *WorkspaceServiceImpl {
	workspaceServiceOnce.Do(func() {
		workspaceService = &WorkspaceServiceImpl{
			workspaceRepository: workspaceRepository,
		}
	})

	return workspaceService
}
