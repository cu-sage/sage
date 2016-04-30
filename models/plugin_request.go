package models

type PluginRequest struct {
	Test *Block  `json:"test"`
	App  *string `json:"app"`
}
