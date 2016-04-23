package models

type PluginRequest struct {
	Test *Block `json:"test"`
	App  *App   `json:"app"`
}
