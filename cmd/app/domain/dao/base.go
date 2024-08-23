package dao

import (
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	CreatedAt time.Time `gorm:"column:created_at;type:TIMESTAMP;null;default:null" json:"created_at"`
	CreatedBy int       `gorm:"column:created_by" json:"created_by"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:TIMESTAMP;null;default:null" json:"updated_at"`
	UpdatedBy int       `gorm:"column:updated_by" json:"updated_by"`
	DeletedAt time.Time `gorm:"column:deleted_at;type:TIMESTAMP;null;default:null" json:"deleted_at"`
	DeletedBy int       `gorm:"column:deleted_by" json:"deleted_by"`
}

// BeforeCreate hook to set CreatedAt and CreatedBy
func (base *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	base.CreatedAt = time.Now()
	base.CreatedBy = -1
	//if userID, ok := tx.Statement.Context.Value("user_id").(int); ok {
	//	base.CreatedBy = userID
	//}
	return
}
