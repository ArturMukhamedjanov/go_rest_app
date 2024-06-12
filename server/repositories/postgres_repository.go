package repositories

import (
	"fmt"
	"go_rest_app/server/configs"
	"log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresRepo struct {
	db *sqlx.DB
}

func (repo *PostgresRepo) InitDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		configs.Configs.Host, configs.Configs.Port, configs.Configs.User, configs.Configs.Password)
	//connStr := "host=localhost port=5432 user=postgres password=postgres sslmode=disable"
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to db: %v", err)
	}
	if !checkDataBaseExists(db, configs.Configs.DBName) {
		log.Printf("Database with name %s does not exist. Creating...\n", configs.Configs.DBName)
		createDatabase(db, configs.Configs.DBName)
	}
	repo.db = db
	log.Fatalf("Successfully connected to DB")
}

func checkDataBaseExists(defaultDB *sqlx.DB, dbName string) bool {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
	err := defaultDB.Get(&exists, query)
	if err != nil {
		log.Fatalf("Error checking if database exists: %v", err)
	}
	return exists
}

func createDatabase(defaultDB *sqlx.DB, dbName string) {
	_, err := defaultDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName))
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
	}
	log.Printf("Database %s created successfully\n", dbName)
}
