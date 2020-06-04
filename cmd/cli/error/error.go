package error

// Exit error
type Exit struct {
	Err     error
	Code    int
	Details string
}

// WrapWithCode wraps error with specified code
func WrapWithCode(err error, code int, details string) *Exit {
	return &Exit{
		Err:     err,
		Code:    code,
		Details: details,
	}
}

// Wrap wraps error with status 1
func Wrap(err error, log string) *Exit {
	return WrapWithCode(err, 1, log)
}

func (e *Exit) Error() string {
	return e.Err.Error()
}
