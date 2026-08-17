package main

import (
	"flag"
	"log"
	"net/http"

	"mountainrescue/internal/rescue"
)

func main() {
	address := flag.String("addr", ":8080", "HTTP listen address")
	flag.Parse()

	repository := rescue.NewMemoryRepository()
	service := rescue.NewService(repository)
	server := rescue.NewHTTPServer(service)

	log.Fatal(http.ListenAndServe(*address, server))
}
