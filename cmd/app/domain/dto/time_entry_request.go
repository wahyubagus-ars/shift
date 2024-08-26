package dto

import "time"

type TimeEntryRequest struct {
	TaskId      string    `json:"task_id"`
	ProjectId   int       `json:"project_id"`
	StartTime   time.Time `json:"start_time"`
	EndTime     time.Time `json:"end_time"`
	Duration    int       `json:"duration"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsBillable  bool      `json:"is_billable"`
}
