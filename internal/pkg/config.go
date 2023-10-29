package pkg

import (
	"github.com/asyauqi1511/test/internal/entity"
	"gopkg.in/yaml.v3"
	"os"
)

func LoadConfig() (entity.AppConfig, error) {

	var config entity.AppConfig

	// Read YAML file.
	yamlFile, err := os.ReadFile("file/config/config.yaml")
	if err != nil {
		return config, err
	}

	// Unmarhal YAML to config.
	err = yaml.Unmarshal(yamlFile, &config)
	return config, err
}
