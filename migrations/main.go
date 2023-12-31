// This is custom goose binary with sqlite3 support only.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"

	_ "migration/migrations"

	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/lib/pq"

	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

var (
	flags = flag.NewFlagSet("goose", flag.ExitOnError)
	dir   = flags.String("dir", ".", "directory with migration files")
)

type MigrationConfig struct {
	dbname   string
	username string
	password string
	port     string
	host     string
}

func getEnv(envName string) string {
	// println("here read ", envName)
	env := os.Getenv(envName)

	if env == "" {
		log.Fatalf("Environement variable not found: %s", envName)
	}

	return env
}

func readEnv() MigrationConfig {

	APP_ENV := getEnv("APP_ENV")
	err := godotenv.Load(fmt.Sprintf(".env.%s", APP_ENV))

	if err != nil {
		log.Fatalf("Error loging file .env.%s %s", APP_ENV, err)
	}

	migrationConfig := MigrationConfig{
		dbname:   getEnv("PG_DB_NAME"),
		username: getEnv("PG_USERNAME"),
		password: getEnv("PG_PASSWORD"),
		host:     getEnv("PG_HOST"),
		port:     getEnv("PG_PORT"),
	}

	return migrationConfig
}

func main() {
	flags.Parse(os.Args[1:])
	args := flags.Args()

	// if len(args) < 2 {
	// 	flags.Usage()
	// 	return
	// }

	config := readEnv()
	dbstring := fmt.Sprintf(
		"postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		config.username,
		config.password,
		config.host,
		config.port,
		config.dbname,
	)

	command := args[0]
	// println(command, dbstring, migrations.testmig())
	db, err := goose.OpenDBWithDriver("postgres", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[2:]...)
	}

	// goose.RunWithOptions()

	// if err := goose.RunContext(command, db, *dir, arguments...); err != nil {
	// 	log.Fatalf("goose %v: %v", command, err)
	// }

	ctx := context.WithValue(context.TODO(), "dbname", config.dbname)

	if err := goose.RunContext(ctx, command, db, "./migrations", arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}

// ➜  go-users ./mig "postgresql://postgres:postgres@0.0.0.0:5432/test?sslmode=disable" up
