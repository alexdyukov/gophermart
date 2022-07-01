package gateway

import "net/http"

// AccrualGateway is struct represents external api which for this service act like DataProvider
// and utilize Gateway pattern. It is here as a temporary mock in case we need to go to Accrual from Gophermart.
type AccrualGateway struct {
	_ http.Client
}
