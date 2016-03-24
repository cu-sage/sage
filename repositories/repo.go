package repositories

import (
	"fmt"

	"bitbucket.org/sage/models"
	"github.com/satori/go.uuid"
)

type Repository struct {
	appsMap map[string]*models.App
	tsMap   map[string]*models.TestSuite
}

func NewRepository() *Repository {
	return &Repository{
		appsMap: make(map[string]*models.App),
		tsMap:   make(map[string]*models.TestSuite),
	}
}

func (r *Repository) SaveApp(app *models.App) error {
	app.ID = uuid.NewV4().String()
	r.appsMap[app.ID] = app

	return nil
}

func (r *Repository) SaveAssessment(ts *models.TestSuite) error {
	ts.ID = uuid.NewV4().String()
	r.tsMap[ts.ID] = ts

	return nil
}

func (r *Repository) GetApp(id string) (models.App, error) {
	app, prs := r.appsMap[id]
	if !prs {
		return models.App{}, fmt.Errorf("Project %q not found", id)
	}

	return *app, nil
}

func (r *Repository) GetAssessment(id string) (models.TestSuite, error) {
	assessment, prs := r.tsMap[id]
	if !prs {
		return models.TestSuite{}, fmt.Errorf("Assessment %q not found", id)
	}

	return *assessment, nil
}
