package web

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/alexdyukov/gophermart/internal/gophermart/domain/usecase"
)

// AccrualGateway is struct represents external api which for this service act like secondary adapter
// and utilize Gateway pattern.
type AccrualGateway struct {
	client http.Client
	addr   string
	path   string
	proto  string
}

func NewAccrualGateway(addr, path string) *AccrualGateway {
	return &AccrualGateway{ // nolint:exhaustivestruct // ok
		addr:  addr,
		path:  path,
		proto: "http://",
	}
}

func (ag *AccrualGateway) GetOrderCalculationState(orderNumber int) (*usecase.CalculationStateDTO, error) {
	numStr := strconv.Itoa(orderNumber)

	log.Println(ag.addr + ag.path + numStr)

	response, err := ag.client.Get(ag.proto + ag.addr + ag.path + numStr)
	if err != nil {
		log.Println(err)

		return nil, err // nolint:wrapcheck // ok
	}

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	dto := usecase.CalculationStateDTO{} // nolint:exhaustivestruct // ok

	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		return nil, err // nolint:wrapcheck // ok
	}

	return &dto, nil
}
