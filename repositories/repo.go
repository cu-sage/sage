package repositories

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"bitbucket.org/sage/models"
	"github.com/satori/go.uuid"
)

type Repository struct {
	session *mgo.Session
}

func NewRepository(session *mgo.Session) *Repository {
	return &Repository{
		session: session,
	}
}

func (r *Repository) SaveApp(app *models.App) error {
	app.ID = uuid.NewV4().String()

	s := r.session.Copy()
	c := s.DB(dbName).C(projectCollection)
	defer s.Close()

	err := c.Insert(app)
	if err != nil {
		log.Printf("Error inserting project: %s\n", err.Error())
	}

	return err
}

func (r *Repository) SaveAssessment(ts *models.TestSuite) error {
	ts.ID = uuid.NewV4().String()

	s := r.session.Copy()
	c := s.DB(dbName).C(assessmentCollection)
	defer s.Close()

	err := c.Insert(ts)
	if err != nil {
		log.Printf("Error inserting assessment: %s\n", err.Error())
	}

	return err
}

func (r *Repository) GetApp(id string) (models.App, error) {
	s := r.session.Copy()
	c := s.DB(dbName).C(projectCollection)
	defer s.Close()

	var result models.App
	err := c.Find(bson.M{"id": id}).One(&result)

	return result, err
}

func (r *Repository) GetAssessment(id string) (models.TestSuite, error) {
	s := r.session.Copy()
	c := s.DB(dbName).C(assessmentCollection)
	defer s.Close()

	var result models.TestSuite
	err := c.Find(bson.M{"id": id}).One(&result)

	return result, err
}
