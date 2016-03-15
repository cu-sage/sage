package models

import "fmt"

type TestSuite struct {
	TestCases []*Block `xml:"block"`
}

func (t *TestSuite) Log() {
	for _, block := range t.TestCases {
		fmt.Println("Block")
		block.Log()
		fmt.Println()
	}
}
