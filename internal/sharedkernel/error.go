package sharedkernel

type AppError struct {
	err error
	// packageName string
	// funcName    string
}

func (e *AppError) Error() string {
	return e.err.Error()
}
