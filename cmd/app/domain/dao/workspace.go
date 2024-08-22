package dao

import (
	"database/sql/driver"
	"encoding/json"
)

type Workspace struct {
	ID             int     `gorm:"column:id;primaryKey;autoIncrement"`
	Name           string  `gorm:"column:name;size:255;not null"`
	Description    string  `gorm:"column:description;type:text"`
	ProfilePicture string  `gorm:"column:profile_picture;size:255"`
	Metadata       JSONMap `gorm:"column:metadata;type:json"`
	BaseModel
}

type JSONMap map[string]interface{}

func (j *JSONMap) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), j)
}

func (j JSONMap) Value() (driver.Value, error) {
	return json.Marshal(j)
}

func (Workspace) TableName() string {
	return "workspace"
}
