package main

import (
	"log"
	"net/http"

	"bitbucket.org/sage/handlers"
	"bitbucket.org/sage/repositories"
    
    "github.com/gorilla/handlers"
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
	r := handler.RegisterHandlers(a)

	log.Printf("Listening on port %s\n", port)
    
    allowedHeaders := []string{"Content-Type"}
    allowedMethods := []string{"POST", "GET", "OPTIONS"}
	log.Fatal(http.ListenAndServe(port,
        handlers.CORS(
            handlers.AllowedMethods(allowedMethods),
            handlers.AllowedHeaders(allowedHeaders),
            )(r)))
}
