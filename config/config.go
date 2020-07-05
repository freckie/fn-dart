package config

import (
	"encoding/json"
	"io/ioutil"

	"fn-dart/models"
)

func LoadConfig(filePath string) (*models.Config, error) {
	cfg := &models.Config{}

	dataBytes, err := ioutil.ReadFile(filePath)
	if err != nil {
		return cfg, err
	}

	json.Unmarshal(dataBytes, cfg)

	return cfg, nil
}
