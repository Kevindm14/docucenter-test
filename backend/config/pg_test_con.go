package config

import (
	"fmt"
	"os"

	"github.com/Kevindm14/docucenter-test/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PGConfigTest struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

var DBTest *gorm.DB

func PgDBConnectionTest() *gorm.DB {
	var err error

	connectionURL := PGConfig{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   "deliveries_db_test",
	}

	conInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=disable",
		connectionURL.Host,
		connectionURL.Port,
		connectionURL.User,
		connectionURL.Password,
	)

	dbName := connectionURL.DBName

	connDbUrl := fmt.Sprintf("%s dbname=%s", conInfo, dbName)
	DB, err = gorm.Open(postgres.Open(conInfo), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	count := 0
	DB.Raw("SELECT count(*) FROM pg_database WHERE datname = ?", dbName).Scan(&count)
	if count == 0 {
		sql := fmt.Sprintf("CREATE DATABASE %s", dbName)
		DB.Exec(sql)
	}

	DB, err = gorm.Open(postgres.Open(connDbUrl), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB.AutoMigrate(
		&models.Customer{},
		&models.GroundDelivery{},
		&models.MaritimeDelivery{},
	)

	fmt.Println("Database connected successfully")

	return DB
}
