package system

type AppError struct {
	Message string
}

func (e *AppError) Error() string {
	return e.Message
}
