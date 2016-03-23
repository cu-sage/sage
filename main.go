package main

import (
	"log"
	"net/http"

	"bitbucket.org/sage/handlers"
	"bitbucket.org/sage/repositories"
)

const (
	port = ":8080"
)

func main() {
	a := &handler.Assessment{
		Repo: &repositories.Repository{},
	}
	handlers := handler.RegisterHandlers(a)

	log.Printf("Listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(port, handlers))
}

// func Foo() {
// 	app, err := parsers.ParseSB2("./data/project.json")
// 	if err != nil {
// 		log.Printf("Error parsing SB2 file: %s", err.Error())
// 	}
//
// 	testSuite, err := parsers.ParseTest("./data/test.xml")
// 	if err != nil {
// 		log.Printf("Error parsing test file: %s", err.Error())
// 	}
//
// 	testRunner := test_runner.NewTestRunner()
// 	for _, test := range testSuite.TestCases {
// 		result, err := testRunner.RunTest(test, app)
// 		if err != nil {
// 			log.Printf("Error running test: %s", err.Error())
// 		}
//
// 		log.Printf("Pass: %+v", result.Pass)
// 		for _, action := range result.Actions {
// 			log.Printf("Action: %s\n", action)
// 		}
// 		log.Println()
// 	}
// }
