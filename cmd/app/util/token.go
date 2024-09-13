package util

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/domain/dto"
	"go-shift/cmd/app/domain/dto/system"
	"strings"
	"time"
)

func GetTokenPayload[T any](jwtClaimDto T, token string) (*T, error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		log.Error("Invalid token format")
		return nil, &system.AppError{
			Message: "Invalid token format",
		}
	}

	// Decode the second part, which is the payload
	payload, err := jwt.DecodeSegment(parts[1])
	if err != nil {
		log.Error("Error decoding payload:", err)
		return nil, err
	}

	err = json.Unmarshal(payload, &jwtClaimDto)
	if err != nil {
		log.Error("Error when unmarshalling payload:", err)
		return nil, nil
	}

	return &jwtClaimDto, nil
}

func CreateAccessToken(secret string, expiry int) (accessToken string, err error) {
	exp := time.Now().Add(time.Hour * time.Duration(expiry)).Unix()
	claims := &dto.JWTClaimsPayloadAccessToken{
		UserId: 13,
		Name:   "Wahyu Bagus",
		Email:  "user.Email",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: exp,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}
	return t, err
}

func CreateRefreshToken(secret string, expiry int) (refreshToken string, err error) {
	claimsRefresh := &dto.JWTRefreshClaimsAccessToken{
		ID:     "1",
		UserId: 1,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(expiry)).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claimsRefresh)
	rt, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return rt, err
}

func IsAuthorized(requestToken string, secret string) (bool, error) {
	_, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return false, err
	}
	return true, nil
}

func ExtractIDFromToken(requestToken string, secret string) (string, error) {
	token, err := jwt.Parse(requestToken, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok && !token.Valid {
		return "", fmt.Errorf("Invalid Token")
	}

	return claims["user_id"].(string), nil
}
