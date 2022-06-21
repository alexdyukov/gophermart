package usecase

type PrettyError interface {
	Pretty() string
}

type ErrorUsecase struct {
	err         error
	pretty      string
	funcName    string
	packageName string
}

func NewErrorUsecase(e error, fn, pn string) ErrorUsecase {
	return ErrorUsecase{
		err:         e,
		funcName:    "",
		packageName: "",
	}
}

func (e ErrorUsecase) SetPretty(msg string) {

}

func (e *ErrorUsecase) Pretty() string {
	return e.pretty
}

func (e *ErrorUsecase) Error() string {
	return e.err.Error()
}
