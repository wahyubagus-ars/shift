package api_service

import (
	"bytes"
	"encoding/json"
	"fmt"
	log "github.com/sirupsen/logrus"
	"go-shift/cmd/app/domain/dto"
	"net/http"
	"sync"
)

var (
	oauthApiService     *OauthApiServiceImpl
	oauthApiServiceOnce sync.Once
)

type OauthApiService interface {
	GetAccessToken(code string, clientId string, clientSecret string, url string, redirectUri string) (dto.AccessTokenDto, error)
}

type OauthApiServiceImpl struct {
}

func (svc *OauthApiServiceImpl) GetAccessToken(code string, clientId string, clientSecret string, url string, redirectUri string) (dto.AccessTokenDto, error) {
	var err error
	var client = &http.Client{}
	var data dto.AccessTokenDto

	body := fmt.Sprintf(`{
		"client_id": "%s",
		"client_secret": "%s",
		"code": "%s",
		"grant_type": "authorization_code",
		"redirect_uri": "%s"
	}`, clientId, clientSecret, code, redirectUri)

	log.Info("body request :: ", body)
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

func ProvideOauthApiService() *OauthApiServiceImpl {
	oauthApiServiceOnce.Do(func() {
		oauthApiService = &OauthApiServiceImpl{}
	})

	return oauthApiService
}
