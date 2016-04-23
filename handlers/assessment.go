package handlers

import (
	"encoding/json"
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"time"

	"bitbucket.org/sage/models"
	"bitbucket.org/sage/parsers"
	"bitbucket.org/sage/repositories"
	"bitbucket.org/sage/runners"
	"bitbucket.org/sage/utils"
)

type Assessment struct {
	Repo   *repositories.Repository
	Runner *runners.TestRunner
}

func NewAssessmentHandler(repo *repositories.Repository, plugins *models.PluginConfig) *Assessment {
	return &Assessment{
		Repo:   repo,
		Runner: runners.NewTestRunner(plugins),
	}
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

	vars := mux.Vars(r)
	testSuite.ID = vars["aid"]

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

	vars := mux.Vars(r)
	app.StudentID = vars["sid"]
	app.AssignmentID = vars["aid"]
	app.TimeSubmitted = time.Now().Unix()

	err = a.Repo.SaveApp(app)
	if err != nil {
		log.Printf("Error saving project: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	utils.WriteJSON(w, http.StatusAccepted, models.PostResult{ID: app.StudentID})
}

func (a *Assessment) GetAssessment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	sid := vars["sid"]
	aid := vars["aid"]

	assignment, err := a.Repo.GetLatestAssignmentFromStudent(sid, aid)
	if err != nil {
		log.Printf("Error retrieving project: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	assessment, err := a.Repo.GetAssessment(aid)
	if err != nil {
		log.Printf("Error retrieving assessment: %s\n", err.Error())
		utils.WriteError(w, err)
		return
	}

	results := []*models.TestResult{}
	for _, test := range assessment.TestCases {
		result, err := a.Runner.RunTest(test, &assignment)
		if err != nil {
			log.Printf("Error running test: %s", err.Error())
		}

		results = append(results, result)
	}

	utils.WriteJSON(w, http.StatusOK, results)
}
