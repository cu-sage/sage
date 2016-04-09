package models

type TestResult struct {
	Pass    bool `json:"pass"`
    Description string `json:"description"`
	Actions []string `json:"actions"`
}
