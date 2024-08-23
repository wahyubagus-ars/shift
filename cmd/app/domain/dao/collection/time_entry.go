package collection

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TimeEntry struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID      int                `bson:"user_id,omitempty" json:"user_id,omitempty"`
	TaskID      primitive.ObjectID `bson:"task_id,omitempty" json:"task_id,omitempty"`
	ProjectID   int                `bson:"project_id,omitempty" json:"project_id,omitempty"`
	StartTime   time.Time          `bson:"start_time,omitempty" json:"start_time,omitempty"`
	EndTime     time.Time          `bson:"end_time,omitempty" json:"end_time,omitempty"`
	Duration    int                `bson:"duration,omitempty" json:"duration,omitempty"`
	Title       string             `bson:"title,omitempty" json:"title,omitempty"`
	Description string             `bson:"description,omitempty" json:"description,omitempty"`
	IsBillable  bool               `bson:"is_billable,omitempty" json:"is_billable,omitempty"`
	BaseModel
}
