package table

type Workspace struct {
	ID             int           `gorm:"column:id;primaryKey;autoIncrement"`
	Name           string        `gorm:"column:name;size:255;not null"`
	Description    string        `gorm:"column:description;type:text"`
	ProfilePicture string        `gorm:"column:profile_picture;size:255"`
	Metadata       JSONMap       `gorm:"column:metadata;type:json"`
	UserAccounts   []UserAccount `gorm:"many2many:user_workspace;foreignKey:ID;joinForeignKey:WorkspaceID;References:ID;joinReferences:UserID"`
	BaseModel
}

func (Workspace) TableName() string {
	return "workspace"
}
