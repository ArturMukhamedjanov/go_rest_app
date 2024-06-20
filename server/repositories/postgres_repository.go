package repositories

import (
	"fmt"
	"go_rest_app/models"
	"go_rest_app/server/configs"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresRepo struct {
	db *gorm.DB
}

func (repo *PostgresRepo) InitDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		configs.Configs.Host, configs.Configs.Port, configs.Configs.User, configs.Configs.Password)
	// Подключаемся к серверу PostgreSQL без указания базы данных
	serverDB, err := gorm.Open(postgres.Open(connStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatalf("Unable to connect to server: %v", err)
	}

	// Проверяем, существует ли база данных
	if !checkDatabaseExists(serverDB, configs.Configs.DBName) {
		log.Printf("Database with name %s does not exist. Creating...\n", configs.Configs.DBName)
		createDatabase(serverDB, configs.Configs.DBName)
	}

	// Подключаемся к созданной базе данных
	dbConnStr := fmt.Sprintf("%s dbname=%s", connStr, configs.Configs.DBName)
	db, err := gorm.Open(postgres.Open(dbConnStr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
		os.Exit(0)
	}
	log.Printf("Successfully connected to DB")
	initTables(db)
	log.Printf("Successfully migrate models")
	repo.db = db

}

func checkDatabaseExists(serverDB *gorm.DB, dbName string) bool {
	var exists bool
	query := fmt.Sprintf("SELECT EXISTS (SELECT 1 FROM pg_database WHERE datname = '%s')", dbName)
	row := serverDB.Raw(query).Row()
	err := row.Scan(&exists)
	if err != nil {
		log.Fatalf("Error checking if database exists: %v", err)
		os.Exit(0)
	}
	return exists
}

func createDatabase(serverDB *gorm.DB, dbName string) {
	err := serverDB.Exec(fmt.Sprintf("CREATE DATABASE %s", dbName)).Error
	if err != nil {
		log.Fatalf("Error creating database: %v", err)
		os.Exit(0)
	}
	log.Printf("Database %s created successfully\n", dbName)
}

func initTables(serverDB *gorm.DB) {
	err := serverDB.AutoMigrate(&models.User{}, &models.Record{})
	if err != nil {
		log.Fatalf("Unable to migrate models to db")
		os.Exit(0)
	}

}

func (repo *PostgresRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := repo.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Printf("User not found: %v", username)
			return nil, nil
		}
		log.Printf("Error fetching user by username %v: %v", username, err)
		return nil, err
	}
	return &user, nil
}

func (repo *PostgresRepo) CreateUser(user models.User) (string, error) {
	if err := repo.db.Create(&user).Error; err != nil {
		return "", err
	}
	return "User created successfully", nil
}

func (repo *PostgresRepo) DeleteUserByUsername(username string) string {
	
	return "User deleted successfully"
}