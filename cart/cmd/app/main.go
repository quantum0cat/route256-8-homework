package main

import (
	"flag"
	"log"
	"net/http"
	itemhandler "route256/cart/internal/app/http/item"
	"route256/cart/internal/pkg/clients/loms"
	itemservice "route256/cart/internal/pkg/services/item"
)

func main() {
	opts := newOptions()
	lomsClient, err := loms.New("loms client", opts.lomsAddr)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/item/add", itemhandler.NewAdd(itemservice.NewAdd(lomsClient)).Handle)
	http.HandleFunc("/item/delete", itemhandler.NewDelete(itemservice.NewDelete()).Handle)
	log.Fatal(http.ListenAndServe(opts.addr, nil))
}

type options struct {
	addr     string
	lomsAddr string
}

func newOptions() *options {
	const (
		defaultAddr     = ":8080"
		defaultLomsAddr = "http://loms:8080"
	)

	result := &options{}
	flag.StringVar(&result.addr, "addr", defaultAddr, "server address, default: "+defaultAddr)
	flag.StringVar(&result.lomsAddr, "loms_addr", defaultLomsAddr, "loms address, default: "+defaultLomsAddr)
	flag.Parse()
	return result
}
