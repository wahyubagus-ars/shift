package service

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/constant"
	"go-shift/cmd/app/domain/dao/table"
	"go-shift/cmd/app/domain/dto/request"
	"go-shift/cmd/app/repository"
	"go-shift/pkg"
	"sync"
)

var (
	workspaceService     *WorkspaceServiceImpl
	workspaceServiceOnce sync.Once
)

type WorkspaceService interface {
	GetWorkspace()
	CreateWorkspace(c *gin.Context)
	//UpdateWorkspace()
	//DeleteWorkspace()
}

type WorkspaceServiceImpl struct {
	workspaceRepository repository.WorkspaceRepository
	userRepository      repository.UserRepository
}

func (svc *WorkspaceServiceImpl) GetWorkspace() {
	svc.workspaceRepository.GetWorkspace()
}

func (svc *WorkspaceServiceImpl) CreateWorkspace(c *gin.Context) {
	defer pkg.PanicHandler(c)

	log.Info("start to execute create new workspace process")

	var createWorkspaceReq request.CreateWorkspace

	err := c.BindJSON(&createWorkspaceReq)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	workspace := &table.Workspace{
		Name:           createWorkspaceReq.Name,
		Description:    createWorkspaceReq.Description,
		ProfilePicture: createWorkspaceReq.ProfilePicture,
	}

	user, err := svc.userRepository.FindUserById(13)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	workspaceData, err := svc.workspaceRepository.AddWorkspace(workspace, user)
	if err != nil {
		pkg.PanicException(constant.UnknownError)
	}

	c.JSON(201, gin.H{"workspace": workspaceData})
}

func ProvideWorkspaceService(workspaceRepository repository.WorkspaceRepository, userRepository repository.UserRepository) *WorkspaceServiceImpl {
	workspaceServiceOnce.Do(func() {
		workspaceService = &WorkspaceServiceImpl{
			workspaceRepository: workspaceRepository,
			userRepository:      userRepository,
		}
	})

	return workspaceService
}
