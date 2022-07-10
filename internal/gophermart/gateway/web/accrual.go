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
	return &AccrualGateway{
		addr:  addr,
		path:  path,
		proto: "http://",
	}
}

func (ag *AccrualGateway) GetOrderCalculationState(orderNumber int64) (*usecase.CalculationStateDTO, error) {
	numStr := strconv.FormatInt(orderNumber, 10)

	log.Println(ag.addr + ag.path + numStr)

	response, err := ag.client.Get(ag.addr + ag.path + numStr)
	if err != nil {
		log.Println(err)

		return nil, err
	}

	defer func() {
		err = response.Body.Close()
		if err != nil {
			log.Println(err)
		}
	}()

	bytes, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	log.Println("income bytes sub service", string(bytes))

	dto := usecase.CalculationStateDTO{}

	err = json.Unmarshal(bytes, &dto)
	if err != nil {
		return nil, err
	}

	log.Println("unmarshaled dto", dto)

	return &dto, nil
}
