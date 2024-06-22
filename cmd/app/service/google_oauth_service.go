package service

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/domain/dto"
	"go-shift/cmd/app/repository"
	"gorm.io/gorm"
	"net/http"
	"os"
	"strings"
	"sync"
)

var (
	googleService     *GoogleOauthServiceImpl
	googleServiceOnce sync.Once
)

type GoogleOauthServiceImpl struct {
	userRepository repository.UserRepository
}

func (svc *GoogleOauthServiceImpl) Login(c *gin.Context) {
	code := c.Request.Header["Authorization-Code"]

	clientId := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	clientSecret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	getTokenUrl := os.Getenv("GOOGLE_ACCESS_TOKEN_URL")

	token, err := GetAccessToken(code[0], clientId, clientSecret, getTokenUrl)
	if err != nil {
		return
	}

	payload, err := getTokenPayload(token.IdToken)

	if err != nil {
		log.Error("Error when try decode token")
		return
	}

	_, err = svc.userRepository.FindUserByEmail(payload.Email)

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			svc.userRepository.SaveInitiateUser(payload.Email, 1)
		}
	}

	// TODO: add process to save token data to db redis

	c.JSON(200, gin.H{
		"response_key": "success",
		"message":      "login api google",
		"data":         token,
	})
}

func GetAccessToken(code string, clientId string, clientSecret string, url string) (dto.AccessTokenDto, error) {
	var err error
	var client = &http.Client{}
	var data dto.AccessTokenDto

	body := fmt.Sprintf(`{
		"client_id": "%s",
		"client_secret": "%s",
		"code": "%s",
		"grant_type": "authorization_code",
		"redirect_uri": "https://localhost:8000"
	}`, clientId, clientSecret, code)

	bodyBytes := []byte(body)

	payload := bytes.NewBuffer(bodyBytes)
	request, err := http.NewRequest("POST", url, payload)
	if err != nil {
		return data, err
	}

	response, err := client.Do(request)
	if err != nil {
		return data, err
	}
	defer response.Body.Close()

	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return data, err
	}

	return data, nil
}

func getTokenPayload(token string) (dto.JWTClaimsPayload, error) {
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
		log.Error("Error unmarshaling payload:", err)
		return dto.JWTClaimsPayload{}, nil
	}

	return claims, nil
}

func ProvideGoogleOauthService(ur repository.UserRepository) *GoogleOauthServiceImpl {
	googleServiceOnce.Do(func() {
		googleService = &GoogleOauthServiceImpl{
			userRepository: ur,
		}
	})

	return googleService
}
