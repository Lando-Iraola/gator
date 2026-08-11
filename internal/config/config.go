package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

func Read() (Config, error) {
	file, err := os.UserHomeDir()
	nilConfig := Config{Db_url: ""}
	if err != nil {

		return nilConfig, err
	}
	configFile, err := os.ReadFile(file + "/" + configFileName)

	if err != nil {
		return nilConfig, err
	}

	var config Config
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		return nilConfig, err
	}

	return config, nil
}

func write(cfg Config) error {
	filePath, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	filePath = filePath + "/" + configFileName
	file, err := os.Create(filePath)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(cfg); err != nil {
		return err
	}
	return nil
}

func SetUser(name string) error {
	config, err := Read()
	if err != nil {
		return err
	}
	config.Current_user_name = name
	err = write(config)
	if err != nil {
		return err
	}
	return nil
}
