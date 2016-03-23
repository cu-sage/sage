package handler

import "github.com/gorilla/mux"

func RegisterHandlers(a *Assessment) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/assessment", a.PostAssessment).
		Methods("POST")
	r.HandleFunc("/project", a.PostProject).
		Methods("POST")

	return r
}
