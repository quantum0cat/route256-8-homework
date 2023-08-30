package item

import (
	"context"
	"log"
)

type DeleteService struct {
}

func NewDelete() *DeleteService {
	return &DeleteService{}
}

func (s DeleteService) Delete(_ context.Context, user int64, sku uint32) error {
	log.Printf("delete item: user: %d, sku: %d", user, sku)
	return nil
}
