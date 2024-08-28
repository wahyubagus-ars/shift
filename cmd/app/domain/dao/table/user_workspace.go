package table

type UserWorkspace struct {
	UserID          int `gorm:"column:user_id;not null"`
	WorkspaceID     int `gorm:"column:workspace_id;not null"`
	WorkspaceRoleID int `gorm:"column:workspace_role_id;not null"`
	BaseModel
}

func (UserWorkspace) TableName() string {
	return "user_workspace"
}
