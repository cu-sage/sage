package repositories

import "bitbucket.org/sage/models"

type Repository struct {
	app *models.App
	ts  *models.TestSuite
}

func (r *Repository) SaveApp(app *models.App) error {
	r.app = app
	return nil
}

func (r *Repository) SaveAssessment(ts *models.TestSuite) error {
	r.ts = ts
	return nil
}
