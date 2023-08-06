package models

import (
	"BookHive/pkg/types"
	"database/sql"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
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

func HashPassword(password string) string {
	var saltRounds int = 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), saltRounds)
	if err != nil {
		panic(err)
	}
	return string(hashedPassword)
}

func MatchPassword(inputPassword string, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}

func ExecSql(db *sql.DB, query string, args ...interface{}) *sql.Rows {
	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}
	return rows
}
