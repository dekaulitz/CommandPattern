package error_handler

type AppException struct {
	s string
}

func NewAppException(s string) *AppException {
	return &AppException{s: s}
}

func (e *AppException) Error() string {
	return e.s
}

