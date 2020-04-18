package config

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func NewFromFile(filePath string) (*File, error) {
	configFile, readFileError := ioutil.ReadFile(filePath)
	if readFileError != nil {
		return nil, readFileError
	}

	var config File
	unmarshalError := yaml.Unmarshal(configFile, &config)
	if unmarshalError != nil {
		return nil, unmarshalError
	}

	return &config, nil
}
