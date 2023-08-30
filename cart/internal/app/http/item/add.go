package item

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"route256/libs/httphandler"
	"time"
)

type AddHandler struct {
	s    Adder
	name string
}

type Adder interface {
	Add(ctx context.Context, user int64, sku uint32, count uint16) error
}

func NewAdd(s Adder) *AddHandler {
	return &AddHandler{s: s, name: "item/add"}
}

type AddRequest struct {
	User  int64  `json:"user,omitempty"`
	Sku   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
}

var (
	ErrIncorrectUser      = errors.New("incorrect user")
	ErrIncorrectSKU       = errors.New("incorrect sku")
	ErrIncorrectItemCount = errors.New("incorrect items count")
)

func (r AddRequest) Validate() error {
	if r.User <= 0 {
		return ErrIncorrectUser
	}
	if r.Sku <= 0 {
		return ErrIncorrectSKU
	}
	if r.Count <= 0 {
		return ErrIncorrectItemCount
	}
	return nil
}

func (h AddHandler) Handle(w http.ResponseWriter, r *http.Request) {
	request := &AddRequest{}

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
	if err := h.s.Add(ctx, request.User, request.Sku, request.Count); err != nil {
		httphandler.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	httphandler.GetSuccessResponse(w)
}
