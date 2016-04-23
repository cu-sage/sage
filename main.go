package main

import (
	"log"
	"net/http"

	"bitbucket.org/sage/handlers"
	"bitbucket.org/sage/repositories"
	"bitbucket.org/sage/utils"

	ghandlers "github.com/gorilla/handlers"
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

	utils.ReadPluginConfig("plugins.json")

	a := handlers.NewAssessmentHandler(
		repositories.NewRepository(session),
	)

	r := handlers.RegisterHandlers(a)

	log.Printf("Listening on port %s\n", port)

	allowedHeaders := []string{"Content-Type"}
	allowedMethods := []string{"POST", "GET", "OPTIONS"}
	log.Fatal(http.ListenAndServe(port,
		ghandlers.CORS(
			ghandlers.AllowedMethods(allowedMethods),
			ghandlers.AllowedHeaders(allowedHeaders),
		)(r)))
}
