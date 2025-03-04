package config

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var DB *gorm.DB

func InitDB() {
	dbUser := GetEnv("DB_USER", "user")
	dbPassword := GetEnv("DB_PASSWORD", "password")
	dbHost := GetEnv("DB_HOST", "db")
	dbPort := GetEnv("DB_PORT", "3306")
	dbName := GetEnv("DB_NAME", "contactsdb")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPassword, dbHost, dbPort, dbName)

	var err error

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to connect to database: %v", err))
	}

	sqlDB, err := DB.DB()
	if err != nil {
		panic("Failed to get database instance!")
	}

	sqlDB.SetMaxOpenConns(10)
	sqlDB.SetMaxIdleConns(5)
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err := DB.AutoMigrate(&models.Contact{}); err != nil {
		panic(fmt.Sprintf("AutoMigrate failed: %v", err))
	}
}
