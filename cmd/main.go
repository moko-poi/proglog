package main

import (
	"log"

	"github.com/moko-poi/proglog/internal/server"
)

func main() {
	srv := server.NewHTTPServer()
	log.Fatal(srv.ListenAndServe())
}
