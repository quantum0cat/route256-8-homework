package services

import (
	"context"
	"log"
)

type StocksService struct {
}

func NewStocks() *StocksService {
	return &StocksService{}
}

func (s StocksService) Stocks(_ context.Context, sku uint32) (uint64, error) {
	log.Printf("stocks: sku: %d", sku)
	return 1000, nil
}
