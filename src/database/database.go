package database

import (
	"log"

	"github.com/psuarezdev/go-api-starter/src/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	dsn := config.GetConnectionString()
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to the database:", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error getting *sql.DB object:", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(0)

	DB = db
}

func GetConnection() *gorm.DB {
	if DB == nil {
		InitDatabase()
	}
	return DB
}
