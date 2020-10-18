package main

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
}

// GetConfig ... Loads configuration file
func GetConfig(params ...string) Configuration {
	f, err := os.Open("dev_config.yaml")
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
