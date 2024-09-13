package constant

type ResponseStatus int
type RedisKey int
type Headers int
type General int

// Constant Api Response
const (
	Success ResponseStatus = iota + 1
	DataNotFound
	UnknownError
	InvalidRequest
	Unauthorized
)

func (r ResponseStatus) GetResponseStatus() string {
	return [...]string{"SUCCESS", "DATA_NOT_FOUND", "UNKNOWN_ERROR", "INVALID_REQUEST", "UNAUTHORIZED"}[r-1]
}

func (r ResponseStatus) GetResponseMessage() string {
	return [...]string{"Success", "Data Not Found", "Unknown Error", "Invalid Request", "Unauthorized"}[r-1]
}

// Constant Redis Key
const (
	UserAccountData RedisKey = iota + 1
	GoogleAccessToken
	GoogleRefreshToken
	GoogleIdToken
	AccessToken
	RefreshToken
)

func (rk RedisKey) GetRedisKey() string {
	return [...]string{"USER_ACCOUNT_DATA", "GOOGLE_ACCESS_TOKEN", "GOOGLE_REFRESH_TOKEN", "GOOGLE_ID_TOKEN", "ACCESS_TOKEN", "REFRESH_TOKEN"}[rk-1]
}
