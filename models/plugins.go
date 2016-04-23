package models

type PluginConfig struct {
	Plugins []Plugin `json:"plugins"`
}

type Plugin struct {
	Type    string `json:"type"`
	Handler string `json:"handler"`
}
