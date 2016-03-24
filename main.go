package main

import (
	"log"
	"net/http"

	"bitbucket.org/sage/handlers"
	"bitbucket.org/sage/repositories"
)

const (
	port = ":8080"
)

func main() {
	a := &handler.Assessment{
		Repo: repositories.NewRepository(),
	}
	handlers := handler.RegisterHandlers(a)

	log.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, handlers))
}
