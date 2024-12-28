package lib

import (
	"embed"
	"encoding/json"
)

//go:embed config/config.json
var configEmbedFS embed.FS

func GetVersion() (int, error) {
	configFileContent, err := configEmbedFS.ReadFile("config/config.json")
	if err != nil {
		return 0, err
	}

	var config struct {
		Version int `json:"version"`
	}
	err = json.Unmarshal(configFileContent, &config)
	if err != nil {
		return 0, err
	}

	return config.Version, nil
}
