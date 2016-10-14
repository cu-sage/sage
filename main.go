package main

import (
	"log"
	"net/http"

	"github.com/cu-sage/sage/handlers"
	"github.com/cu-sage/sage/repositories"
	"github.com/cu-sage/sage/utils"

	ghandlers "github.com/gorilla/handlers"
)

const (
	port = ":8081"
)

func main() {
	session, err := repositories.GetMongoSession()
	if err != nil {
		log.Panicf("Cannot connect to database. Exiting...")
	}
	defer session.Close()

	a := handlers.NewAssessmentHandler(
		repositories.NewRepository(session),
		utils.ReadPluginConfig("plugins.json"),
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
