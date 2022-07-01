package usecase

// UserInputDTO struct is in separate file due to used by both usecases.
type UserInputDTO struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
