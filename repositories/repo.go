package repositories

import (
	"log"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"

	"github.com/cu-sage/sage/models"
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
	s := r.session.Copy()
	c := s.DB(dbName).C(assessmentCollection)
	defer s.Close()

	err := c.Insert(ts)
	if err != nil {
		log.Printf("Error inserting assessment: %s\n", err.Error())
	}

	return err
}

func (r *Repository) GetLatestAssignmentFromStudent(sid, aid string) (models.App, error) {
	s := r.session.Copy()
	c := s.DB(dbName).C(projectCollection)
	defer s.Close()

    log.Printf("studentid: %s, assignmentid: %s\n", sid, aid)

	var result models.App
	err := c.Find(bson.M{"studentid": sid, "assignmentid": aid}).Sort("-timesubmitted").One(&result)

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
