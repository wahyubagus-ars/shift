package dto

import "github.com/dgrijalva/jwt-go"

type AccessTokenDto struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int32  `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	Scope        string `json:"scope"`
	TokenType    string `json:"token_type"`
	IdToken      string `json:"id_token"`

	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}

type JWTClaimsPayloadGoogle struct {
	Iss           string `json:"iss"`
	Azp           string `json:"azp"`
	Aud           string `json:"aud"`
	Sub           string `json:"sub"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	AtHash        string `json:"at_hash"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Iat           int64  `json:"iat"`
	Exp           int64  `json:"exp"`
}

type JWTClaimsPayloadAccessToken struct {
	UserId int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

type JWTRefreshClaimsAccessToken struct {
	ID     string `json:"id"`
	UserId int    `json:"user_id"`
	jwt.StandardClaims
}

type Env struct {
	AccessTokenExpiryHour  int
	RefreshTokenExpiryHour int
	AccessTokenSecret      string
	RefreshTokenSecret     string
}
