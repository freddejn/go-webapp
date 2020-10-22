package config

import (
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

// Configuration ... Struct for configuration
type Configuration struct {
	Server struct {
		PORT string `yaml:"port"`
		Host string `yaml:"host"`
	}
	Resources struct {
		Templates string `yaml:"templates"`
	}
}

// GetConfig ... Loads configuration file
func GetConfig(params ...string) Configuration {
	env := "dev"
	if len(params) > 0 {
		env = params[0]
	}
	f, err := os.Open(env + "_config.yaml")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var cfg Configuration
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		log.Fatal(err)
	}
	return cfg
}
