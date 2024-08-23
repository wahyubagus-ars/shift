package dao

type UserProfile struct {
	ID            int     `gorm:"primaryKey;column:id" json:"id"`
	UserAccountID int     `gorm:"column:user_account_id" json:"user_account_id"`
	FirstName     string  `gorm:"column:first_name" json:"first_name"`
	LastName      string  `gorm:"column:last_name" json:"last_name"`
	DOB           int     `gorm:"column:dob" json:"dob"` // Consider using a date/time type if applicable
	POB           string  `gorm:"column:pob" json:"pob"`
	Status        int     `gorm:"column:status" json:"status"`
	Metadata      JSONMap `gorm:"column:metadata" json:"metadata"`
	BaseModel
}
