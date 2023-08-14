package models

import (
	"BookHive/pkg/utils"
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func dsn() (string, error) {
	config, err := utils.LoadConfig()
	if err != nil {
		return "", err
	}

	// Returning the dsn string
	return fmt.Sprintf("%s:%s@tcp(%s)/%s",
		config.DbUser,
		config.Password,
		config.Host,
		config.DbName,
	), nil
}

func Connection() (*sql.DB, error) {
	// Getting dsn
	dsn, err := dsn()
	if err != nil {
		return nil, err
	}

	// Opening the connection
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error: '%s' when opening DB", err)
		return nil, err
	}

	// Configuring the connection
	db.SetMaxOpenConns(20)
	db.SetMaxIdleConns(20)
	db.SetConnMaxLifetime(time.Minute * 5)

	ctx, cancelfunc := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancelfunc()
	err = db.PingContext(ctx)
	if err != nil {
		log.Fatalf("Error: '%s' occured while pinging DB", err)
		return nil, err
	}
	return db, err
}
