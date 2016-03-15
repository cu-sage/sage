package parsers

import (
	"encoding/xml"
	"io/ioutil"

	"bitbucket.org/sage/models"
)

func ParseTest(path string) (*models.TestSuite, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var testSuite models.TestSuite
	err = xml.Unmarshal(file, &testSuite)
	if err != nil {
		return nil, err
	}

	testSuite.Log()
	return &testSuite, nil
}
