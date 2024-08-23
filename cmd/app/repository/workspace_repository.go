package repository

import (
	"go-shift/cmd/app/domain/dao/table"
	"gorm.io/gorm"
	"sync"
)

var (
	workspaceRepository    *WorkspaceRepositoryImpl
	workspaceRepositoyOnce sync.Once
)

type WorkspaceRepository interface {
	GetWorkspace() ([]table.Workspace, error)
	AddWorkspace(workspace table.Workspace) ([]table.Workspace, error)
}

type WorkspaceRepositoryImpl struct {
	mysql *gorm.DB
}

func (wr *WorkspaceRepositoryImpl) GetWorkspace() ([]table.Workspace, error) {
	return nil, nil
}

func (wr *WorkspaceRepositoryImpl) AddWorkspace(workspace table.Workspace) ([]table.Workspace, error) {
	return nil, nil
}

func ProvideWorkspaceRepository(mysql *gorm.DB) *WorkspaceRepositoryImpl {
	workspaceRepositoyOnce.Do(func() {
		workspaceRepository = &WorkspaceRepositoryImpl{
			mysql: mysql,
		}
	})

	return workspaceRepository
}
