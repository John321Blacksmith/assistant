// Package config provides the tool
// for extraction the settings from
// the json file.
package config

import (
	"encoding/json"
	"os"
	"osint_agent/services/web_agent/internal/domain"
)

// Function LoadConfig retrieves settings data
// from the specified json file and returns a
// prepared Config object.
func LoadConfig(cfgPath string) (*domain.Config, error) {
	cfg := new(domain.Config)

	bytes, err := os.ReadFile(cfgPath)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bytes, cfg)
	return cfg, err
}
