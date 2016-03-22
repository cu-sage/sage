package main

import (
	"log"

	"bitbucket.org/sage/parsers"
	"bitbucket.org/sage/test_runners"
)

func main() {
	app, err := parsers.ParseSB2("./data/project.json")
	if err != nil {
		log.Printf("Error parsing SB2 file: %s", err.Error())
	}

	testSuite, err := parsers.ParseTest("./data/test.xml")
	if err != nil {
		log.Printf("Error parsing test file: %s", err.Error())
	}

	testRunner := test_runner.NewTestRunner()
	for _, test := range testSuite.TestCases {
		result, err := testRunner.RunTest(test, app)
		if err != nil {
			log.Printf("Error running test: %s", err.Error())
		}

		log.Printf("Pass: %+v", result.Pass)
		for _, action := range result.Actions {
			log.Printf("Action: %s\n", action)
		}
		log.Println()
	}
}
