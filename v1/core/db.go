package core

import (
	"database/sql"
	"fmt"

	"github.com/tkanos/gonfig"
)

var (
	host     = ""
	port     = 5432
	user     = ""
	password = ""
	dbName   = ""
)

type configuration struct {
	DB_USERNAME string
	DB_PASSWORD string
	DB_PORT     string
	DB_HOST     string
	DB_NAME     string
}

func getConfig(params ...string) configuration {
	configuration := configuration{}
	env := "dev"

	if len(params) > 0 {
		env = params[0]
	}

	fileName := fmt.Sprintf("./%s_config.json", env)
	gonfig.GetConf(fileName, &configuration)
	return configuration
}

func connect() (*sql.DB, error) {
	configuration := getConfig()

	user = configuration.DB_USERNAME
	password = configuration.DB_PASSWORD
	dbName = configuration.DB_NAME

	var connectionString string = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbName)
	db, err := sql.Open("postgres", connectionString)

	if err != nil {
		return nil, err
	}

	return db, nil
}
