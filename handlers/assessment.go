package handler

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"bitbucket.org/sage/models"
	"bitbucket.org/sage/parsers"
	"bitbucket.org/sage/repositories"
	"bitbucket.org/sage/test_runners"
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

	utils.WriteJSON(w, http.StatusAccepted, models.PostResult{ID: testSuite.ID})
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

	utils.WriteJSON(w, http.StatusAccepted, models.PostResult{ID: app.ID})
}

func (a *Assessment) GetAssessment(w http.ResponseWriter, r *http.Request) {
	projectID := mux.Vars(r)["projectID"]
	assessmentID := mux.Vars(r)["assessmentID"]

	project, err := a.Repo.GetApp(projectID)
	if err != nil {
		log.Printf("Error retrieving project: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	assessment, err := a.Repo.GetAssessment(assessmentID)
	if err != nil {
		log.Printf("Error retrieving assessment: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	testRunner := test_runner.NewTestRunner()
	results := []*models.TestResult{}
	for _, test := range assessment.TestCases {
		result, err := testRunner.RunTest(test, &project)
		if err != nil {
			log.Printf("Error running test: %s", err.Error())
		}

		results = append(results, result)
	}

	utils.WriteJSON(w, http.StatusOK, results)
}