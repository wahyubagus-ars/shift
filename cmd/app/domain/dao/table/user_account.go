package table

import "time"

type UserAccount struct {
	ID               int       `gorm:"primaryKey" json:"id"`
	Email            string    `gorm:"unique;not null" json:"email"`
	EmailVerifiedAt  time.Time `json:"email_verified_at" gorm:"type:TIMESTAMP;null;default:null"`
	Password         string    `json:"password"`
	AuthenticationID int       `gorm:"not null" json:"authentication_id"`
	BaseModel
}

func (UserAccount) TableName() string {
	return "user_account"
}
