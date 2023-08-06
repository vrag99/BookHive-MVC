package utils

import (
	"BookHive/pkg/types"
	"database/sql"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

func LoadConfig() (types.YamlConfig, error) {
	// Opening the config file
	creds, err := os.Open("config.yaml")
	if err != nil {
		fmt.Printf("Error: '%s' while trying to open config.yaml\n", err)
		return types.YamlConfig{}, err
	}
	defer creds.Close()

	// Loading the details in "config"
	var config types.YamlConfig
	decoder := yaml.NewDecoder(creds)

	if err := decoder.Decode(&config); err != nil {
		fmt.Printf("Error: '%s' while decoding config.yaml\n", err)
		return types.YamlConfig{}, err
	}

	return config, nil
}

func ExecSql(db *sql.DB, query string, args ...interface{}) *sql.Rows {
	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	return rows
}
