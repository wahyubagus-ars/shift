package collection

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type EmailVerification struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	UserID           int                `bson:"user_id,omitempty" json:"user_id,omitempty"`
	VerificationType string             `bson:"invitation_type,omitempty"`
	Email            string             `bson:"email,omitempty"`
	Status           string             `bson:"status,omitempty"`
	ExpiredAt        time.Time          `bson:"expired_at,omitempty"`
}
