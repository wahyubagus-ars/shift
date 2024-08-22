package repository

import (
	"go-shift/cmd/app/domain/dao"
	"gorm.io/gorm"
	"sync"
)

var (
	workspaceRepository    *WorkspaceRepositoryImpl
	workspaceRepositoyOnce sync.Once
)

type WorkspaceRepository interface {
	GetWorkspace() ([]dao.Workspace, error)
	AddWorkspace(workspace dao.Workspace) ([]dao.Workspace, error)
}

type WorkspaceRepositoryImpl struct {
	mysql *gorm.DB
}

func (wr *WorkspaceRepositoryImpl) GetWorkspace() ([]dao.Workspace, error) {
	return nil, nil
}

func (wr *WorkspaceRepositoryImpl) AddWorkspace(workspace dao.Workspace) ([]dao.Workspace, error) {
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
