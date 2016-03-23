package utils

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/sage/models"
)

// WriteJSON writes an HTTP response with JSON
func WriteJSON(w http.ResponseWriter, status int, v interface{}) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)

	err := json.NewEncoder(w).Encode(v)
	if err != nil {
		log.Printf("Error encoding JSON: %s", err.Error())
	}
}

// WriteError writes an HTTP response with an error
func WriteError(w http.ResponseWriter, e error) {
	r := &models.Error{
		Message: e.Error(),
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusInternalServerError)

	err := json.NewEncoder(w).Encode(r)
	if err != nil {
		log.Printf("Error encoding error response: %s", err.Error())
	}
}
