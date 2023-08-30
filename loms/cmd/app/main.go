package main

import (
	"flag"
	"log"
	"net/http"
	stockshandler "route256/loms/internal/app/http"
	"route256/loms/internal/pkg/services"
)

func main() {
	opts := newOptions()
	http.HandleFunc("/stocks", stockshandler.NewStocksHandler(services.NewStocks()).Handle)
	log.Fatal(http.ListenAndServe(opts.addr, nil))
}

type options struct {
	addr string
}

func newOptions() *options {
	const defaultAddr = ":8080"

	result := &options{}
	flag.StringVar(&result.addr, "addr", defaultAddr, "server address, default: "+defaultAddr)
	flag.Parse()
	return result
}
