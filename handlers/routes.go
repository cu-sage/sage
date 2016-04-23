package handlers

import "github.com/gorilla/mux"

func RegisterHandlers(a *Assessment) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/assessments/{aid}", a.PostAssessment).
		Methods("POST")
	r.HandleFunc("/students/{sid}/assignments/{aid}", a.PostProject).
		Methods("POST")
	r.HandleFunc("/students/{sid}/assessments/{aid}/results", a.GetAssessment).
		Methods("GET")

	return r
}
