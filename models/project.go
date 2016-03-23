package models

type SB2Project struct {
	Children []struct {
		ObjName string          `json:"objName"`
		Scripts [][]interface{} `json:"scripts"`
	} `json:"children"`
}
