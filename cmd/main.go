package main

import (
	"log"
	"net/http"

	"github.com/VashUber/go-crud/internal/server"
)

func main() {
	s := server.New()

	log.Fatal(http.ListenAndServe(":8080", s))
}
