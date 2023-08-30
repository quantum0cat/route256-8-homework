package http

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"route256/libs/httphandler"
	"time"
)

type StocksRequest struct {
	SKU uint32 `json:"sku,omitempty"`
}

type StocksResponse struct {
	Count uint64 `json:"count,omitempty"`
}

type StocksHandler struct {
	s    StocksService
	name string
}

type StocksService interface {
	Stocks(ctx context.Context, sku uint32) (uint64, error)
}

func NewStocksHandler(s StocksService) *StocksHandler {
	return &StocksHandler{s: s, name: "stocks"}
}

var (
	ErrIncorrectSKU = errors.New("incorrect sku")
)

func (r StocksRequest) Validate() error {
	if r.SKU <= 0 {
		return ErrIncorrectSKU
	}
	return nil
}

func (h StocksHandler) Handle(w http.ResponseWriter, r *http.Request) {
	request := &StocksRequest{}

	if err := json.NewDecoder(r.Body).Decode(request); err != nil {
		httphandler.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}
	if err := request.Validate(); err != nil {
		httphandler.GetErrorResponse(w, h.name, err, http.StatusBadRequest)
		return
	}
	ctx, cancel := context.WithTimeout(r.Context(), time.Second)
	defer cancel()

	count, err := h.s.Stocks(ctx, request.SKU)
	if err != nil {
		httphandler.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	response := &StocksResponse{Count: count}
	raw, err := json.Marshal(response)
	if err != nil {
		httphandler.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	httphandler.GetSuccessResponseWithBody(w, raw)
}
