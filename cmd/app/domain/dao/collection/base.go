package collection

import (
	"time"
)

type BaseModel struct {
	CreatedAt *time.Time `bson:"created_at,omitempty" json:"created_at,omitempty"`
	CreatedBy *int       `bson:"created_by,omitempty" json:"created_by,omitempty"`
	UpdatedAt *time.Time `bson:"updated_at,omitempty" json:"updated_at,omitempty"`
	UpdatedBy *int       `bson:"updated_by,omitempty" json:"updated_by,omitempty"`
	DeletedAt *time.Time `bson:"deleted_at,omitempty" json:"deleted_at,omitempty"`
	DeletedBy *int       `bson:"deleted_by,omitempty" json:"deleted_by,omitempty"`
}
