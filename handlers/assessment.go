package handler

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"bitbucket.org/sage/models"
	"bitbucket.org/sage/parsers"
	"bitbucket.org/sage/repositories"
	"bitbucket.org/sage/utils"
)

type Assessment struct {
	Repo *repositories.Repository
}

func (a *Assessment) PostAssessment(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, readLimit))
	if err != nil {
		log.Printf("Error reading request body: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	var testSuite models.TestSuite
	err = xml.Unmarshal(body, &testSuite)
	if err != nil {
		log.Printf("Error unmarshalling request body: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	err = a.Repo.SaveAssessment(&testSuite)
	if err != nil {
		log.Printf("Error saving project: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}

func (a *Assessment) PostProject(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, readLimit))
	if err != nil {
		log.Printf("Error reading request body: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	var project models.SB2Project
	err = json.Unmarshal(body, &project)
	if err != nil {
		log.Printf("Error unmarshalling request body: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	app, err := parsers.ParseSB2(project)
	if err != nil {
		log.Printf("Error parsing project: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	err = a.Repo.SaveApp(app)
	if err != nil {
		log.Printf("Error saving project: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	w.WriteHeader(http.StatusAccepted)
}
