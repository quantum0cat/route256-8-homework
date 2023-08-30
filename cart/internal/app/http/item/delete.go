package item

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"route256/libs/httphandler"
)

type DeleteHandler struct {
	s    Deleter
	name string
}

type Deleter interface {
	Delete(ctx context.Context, user int64, sku uint32) error
}

func NewDelete(s Deleter) *DeleteHandler {
	return &DeleteHandler{s: s, name: "item/delete"}
}

type DeleteRequest struct {
	User  int64  `json:"user,omitempty"`
	Sku   uint32 `json:"sku,omitempty"`
	Count uint16 `json:"count,omitempty"`
}

func (r DeleteRequest) Validate() error {
	if r.User <= 0 {
		return ErrIncorrectUser
	}
	if r.Sku <= 0 {
		return ErrIncorrectSKU
	}
	return nil
}

func (h DeleteHandler) Handle(w http.ResponseWriter, r *http.Request) {
	request := &DeleteRequest{}

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
	if err := h.s.Delete(ctx, request.User, request.Sku); err != nil {
		httphandler.GetErrorResponse(w, h.name, err, http.StatusInternalServerError)
		return
	}

	httphandler.GetSuccessResponse(w)
}
