package item

import (
	"context"
	"errors"
	"log"
)

type AddService struct {
	stocksProvider StocksProvider
}

type StocksProvider interface {
	GetStocks(ctx context.Context, sku uint32) (uint64, error)
}

func NewAdd(stocksProvider StocksProvider) *AddService {
	return &AddService{
		stocksProvider: stocksProvider,
	}
}

func (s AddService) Add(ctx context.Context, user int64, sku uint32, count uint16) error {
	stocks, err := s.stocksProvider.GetStocks(ctx, sku)
	if err != nil {
		return err
	}
	if stocks < uint64(count) {
		return errors.New("insufficient stocks")
	}
	log.Printf("item add: user: %d, sku: %d, count: %d", user, sku, count)
	return nil
}
