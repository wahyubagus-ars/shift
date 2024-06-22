package dao

type Authentication struct {
	ID      int    `gorm:"primaryKey;column:id" json:"id"`
	Name    string `gorm:"column:name;size:25;not null" json:"name"`
	Enabled bool   `gorm:"column:enabled;default:true" json:"enabled"`
	BaseModel
}
