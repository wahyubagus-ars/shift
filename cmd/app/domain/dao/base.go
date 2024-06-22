package dao

import "time"

type BaseModel struct {
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;null;default:null" json:"created_at"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;null;default:null" json:"updated_at"`
	UpdatedBy int       `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:TIMESTAMP;null;default:null" json:"deleted_at"`
	DeletedBy int       `gorm:"column:deleted_by" json:"deleted_by"`
}
