package table

import "time"

type AuthToken struct {
	ID            int       `gorm:"primaryKey;column:id" json:"id"`
	UserAccountID int       `gorm:"column:user_account_id" json:"user_account_id"`
	AccessToken   string    `gorm:"column:access_token" json:"access_token"`
	RefreshToken  string    `gorm:"column:refresh_token" json:"refresh_token"`
	ExpiresIn     time.Time `gorm:"column:expires_in" json:"expires_in"`
	IsActive      bool      `gorm:"column:is_active;default:false" json:"is_active"`
	BaseModel
}

func (AuthToken) TableName() string {
	return "auth_token"
}
