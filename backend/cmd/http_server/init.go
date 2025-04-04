package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

const (
	path              = ".env"
	pgPortEnvName     = "PG_PORT"
	pgUserEnvName     = "PG_USER"
	pgPasswordEnvName = "PG_PASSWORD"
	pgDatabaseEnvName = "PG_DB"
	pgHostEnvName     = "DB_HOST"
)

func LoadEnv() error {
	err := godotenv.Load(path)
	if err != nil {
		return err
	}
	return nil
}

func DSN() (string, error) {
	host := os.Getenv(pgHostEnvName)
	if len(host) == 0 {
		return "", errors.New("no DB_HOST env value")
	}
	port := os.Getenv(pgPortEnvName)
	if len(host) == 0 {
		return "", errors.New("no PG_PORT env value")
	}
	user := os.Getenv(pgUserEnvName)
	if len(host) == 0 {
		return "", errors.New("no PG_USER env value")
	}
	password := os.Getenv(pgPasswordEnvName)
	if len(host) == 0 {
		return "", errors.New("no PG_PASSWORD env value")
	}
	dbname := os.Getenv(pgDatabaseEnvName)
	if len(host) == 0 {
		return "", errors.New("no PG_DATABASE env value")
	}

	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		host,
		port,
		user,
		password,
		dbname,
	), nil
}
