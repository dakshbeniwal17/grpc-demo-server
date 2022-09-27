package db

import (
	"database/sql"
	"fmt"
	"grpc-demo-server/utils/helpers"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

func PostgresConnect() (*sql.DB, error) {
	config, err := helpers.LoadEnvFile()
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] loading toml file: %+v\n", err)
		return nil, err
	}

	host := config.Get("postgres.host").(string)
	port := config.Get("postgres.port").(int64)
	user := config.Get("postgres.user").(string)
	password := config.Get("postgres.password").(string)
	dbName := config.Get("postgres.dbName").(string)

	if dbHostEnv, lookup := os.LookupEnv("DB_HOST"); lookup {
		host = dbHostEnv
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] opening postgres connection: %+v\n", err)
		return nil, err
	} else {
		fmt.Printf("Connected to database: %s\n", dbName)
	}

	if err := db.Ping(); err != nil {
		fmt.Fprintf(os.Stderr, "[ERROR] pinging db: %+v\n", err)
		return nil, err
	}

	boil.DebugMode = true
	return db, nil
}

func RunDBMigration(migrationURL string, db *sql.DB) {
	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("cannot create postgres driver:", err)
	}
	migration, err := migrate.NewWithDatabaseInstance(migrationURL, "postgres", driver)
	if err != nil {
		log.Fatal("cannot create new migrate instance:", err)
	}

	if err = migration.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("failed to run migrate up:", err)
	}

	log.Println("db migrated successfully")
}
