// This is custom goose binary with sqlite3 support only.

package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

type AuthConfig struct {
	JwtLoginKey string
}

type PGConfig struct {
	Dbname   string
	Username string
	Password string
	Port     string
	Host     string
}

type Config struct {
	PgConfig   PGConfig
	AuthConfig AuthConfig
	Port       string
	ENV        string
}

type AppCtx struct {
	ORM *gorm.DB
}

func getEnv(envKey string) string {
	env := os.Getenv(envKey)

	if env == "" {
		log.Fatalf("Environement variable not found: %s", envKey)
	}

	return env
}

func ReadConfig() Config {
	ENV := getEnv("APP_ENV")

	log.Printf("\nLoading Environment config: %s", ENV)

	if !(ENV == "development" || ENV == "staging") {
		log.Fatalf("Invalid environment input: %s", ENV)
	}

	err := godotenv.Load(fmt.Sprintf(".env.%s", ENV))

	if err != nil {
		log.Fatalf("Error loging file .env %s", err)
	}

	pgConfig := PGConfig{
		Dbname:   getEnv("PG_DB_NAME"),
		Username: getEnv("PG_USERNAME"),
		Password: getEnv("PG_PASSWORD"),
		Host:     getEnv("PG_HOST"),
		Port:     getEnv("PG_PORT"),
	}

	authConfig := AuthConfig{
		JwtLoginKey: getEnv("JWT_LOGIN_KEY"),
	}

	config := Config{
		PgConfig:   pgConfig,
		AuthConfig: authConfig,
		Port:       getEnv("PORT"),
		ENV:        ENV,
	}

	return config
}
