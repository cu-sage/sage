package models

type TestResult struct {
	Pass        bool     `json:"pass"`
	Description string   `json:"description"`
	Actions     []Action `json:"actions"`
}

type Action struct {
	Type    string `json:"type"`
	Command string `json:"command"`
}
