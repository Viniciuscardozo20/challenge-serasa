package app

import (
	"challenge-serasa/api/database"
	. "challenge-serasa/api/helper_tests/h_database"
	"encoding/json"
	"os"
)

type Config struct {
	Passphrase   string   `json:"passphrase"`
	Key          string   `json:"key"`
	MainframeUrl string   `json:"mainframeUrl"`
	Port         int      `json:"port"`
	Database     Database `json:"database"`
}

type Database struct {
	Config                 database.Config `json:"config"`
	NegativationCollection string          `json:"negativation_collection"`
}

func NewConfigFile(filename string) error {
	err := generateConfigFile(filename, configSample())
	if err != nil {
		return err
	}
	return nil
}

func generateConfigFile(filename string, config Config) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	_, err = file.Write(data)
	if err != nil {
		return err
	}
	return nil
}

func configSample() Config {
	return Config{
		Passphrase:   "password",
		Key:          "secretkey",
		MainframeUrl: "http://mainframe.service.com.br:5000",
		Port:         9000,
		Database: Database{
			Config: database.Config{
				Host:     "http://mongo.service.com.br",
				Port:     DBPortTest,
				User:     DBUserTest,
				Password: DBPassTest,
				Database: DBNameTest,
			},
			NegativationCollection: "negativation-collection",
		},
	}
}
