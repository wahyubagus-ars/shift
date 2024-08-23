package util

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/domain/dto"
	"strings"
)

func GetTokenPayload(token string) (dto.JWTClaimsPayload, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		log.Error("Invalid token format")
		return dto.JWTClaimsPayload{}, &dto.AppError{
			Message: "Invalid token format",
		}
	}

	// Decode the second part, which is the payload
	payload, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		log.Error("Error decoding payload:", err)
		return dto.JWTClaimsPayload{}, err
	}

	var claims dto.JWTClaimsPayload
	err = json.Unmarshal(payload, &claims)
	if err != nil {
		log.Error("Error when unmarshalling payload:", err)
		return dto.JWTClaimsPayload{}, nil
	}

	return claims, nil
}
