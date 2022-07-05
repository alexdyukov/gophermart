package web

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/usecase"
)

// AccrualGateway is struct represents external api which for this service act like DataProvider
// and utilize ServiceGateway pattern. It is here as a temporary mock in case we need to go to Accrual from Gophermart.
type AccrualGateway struct {
	client http.Client
	addr   string
	path   string
}

func NewAccrualGateway(addr, path string) *AccrualGateway {
	return &AccrualGateway{ // nolint:exhaustivestruct // ok
		addr: addr,
		path: path,
	}
}

func (ag *AccrualGateway) GetOrderCalculationState(orderNumber int) (*usecase.RegisterOrderCalculationStateDTO, error) {

	numStr := strconv.Itoa(orderNumber)

	fmt.Println(ag.addr + ag.path + numStr)

	response, err := ag.client.Get("http://" + ag.addr + ag.path + numStr)
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	dto := usecase.RegisterOrderCalculationStateDTO{} // nolint:exhaustivestruct // ok

	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	return &dto, nil
}
