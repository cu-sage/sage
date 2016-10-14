package utils

import (
	"encoding/json"
	"log"
	"os"

	"github.com/cu-sage/sage/models"
)

// ReadPluginConfig reads a config file at the given path
// and parses it as a config file for plugins
func ReadPluginConfig(path string) *models.PluginConfig {
	pluginConfig := &models.PluginConfig{}

	file, err := os.Open(path)
	if err != nil {
		log.Printf("Unable to read plugin config at %q: %s\n", path, err.Error())
		return pluginConfig
	}
	defer file.Close()

	jsonParser := json.NewDecoder(file)
	err = jsonParser.Decode(pluginConfig)
	if err != nil {
		log.Printf("Unable to decode plugin config at %q: %s\n", path, err.Error())
	}

	return pluginConfig
}
