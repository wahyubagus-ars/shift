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
	AccessToken
	RefreshToken
	IdToken
)

func (rk RedisKey) GetRedisKey() string {
	return [...]string{"USER_ACCOUNT_DATA", "ACCESS_TOKEN", "REFRESH_TOKEN", "ID_TOKEN"}[rk-1]
}
