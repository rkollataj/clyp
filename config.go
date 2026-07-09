package main

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Comment      string `json:"comment"`
	CloseOnCopy  bool   `json:"close_on_copy"`
	SimplifiedUI bool   `json:"simplified_ui"`
	file         string `json:"-"`
}

func (config *Config) init() {
	config.file = app.configDir + "/config.json"
	if config.isExist() {
		config.load()
	} else {
		config.setDefaults()
		config.save()
	}
}

func (config *Config) isExist() bool {
	_, err := os.Stat(config.file)
	return err == nil
}

func (config *Config) setDefaults() {
	config.Comment = "Documentation: " + app.helpURL
	config.CloseOnCopy = false
	config.SimplifiedUI = false
}

func (config *Config) save() {
	configData := &Config{
		Comment:      config.Comment,
		CloseOnCopy:  config.CloseOnCopy,
		SimplifiedUI: config.SimplifiedUI,
	}
	configJSON, err := json.MarshalIndent(configData, "", "  ")
	if err != nil {
		log.Printf("Failed to marshal config file: %v", err)
		return
	}
	err = os.WriteFile(config.file, configJSON, 0644)
	if err != nil {
		log.Printf("Failed to write config file: %v", err)
		return
	}
}

func (config *Config) load() {
	configJSON, err := os.ReadFile(config.file)
	if err != nil {
		log.Printf("Failed to read config file: %v", err)
		return
	}
	err = json.Unmarshal(configJSON, config)
	if err != nil {
		log.Printf("Failed to parse config file: %v", err)
		return
	}
}
