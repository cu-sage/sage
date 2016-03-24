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
	session, err := repositories.GetMongoSession()
	if err != nil {
		log.Panicf("Cannot connect to database. Exiting...")
	}
	defer session.Close()

	a := &handler.Assessment{
		Repo: repositories.NewRepository(session),
	}
	handlers := handler.RegisterHandlers(a)

	log.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, handlers))
}
