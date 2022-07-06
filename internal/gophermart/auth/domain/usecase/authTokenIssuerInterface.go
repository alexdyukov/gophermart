package usecase

// AuthTokenIssuerGateway interface is in separate file due to used by both usecases.
type AuthTokenIssuerGateway interface {
	IssueWithLoginAndID(login, id string) (string, error)
}
