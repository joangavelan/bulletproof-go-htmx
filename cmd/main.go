package main

import (
	"log"

	"github.com/joangavelan/bulletproof-go-htmx/handlers"
)

func main() {
	server := handlers.GetServer()

	log.Fatal(server.Start())
}