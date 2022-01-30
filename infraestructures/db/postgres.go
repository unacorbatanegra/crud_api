package db

import (
	"fmt"
	"log"

	"github.com/unacorbatanegra/crud_api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresDB interface {
	DB() *gorm.DB
}

type postgresDB struct {
	db *gorm.DB
}

func NewPostgresClient() PostgresDB {

	var err error
	var db *gorm.DB

	c := config.Get()

	dbHost := c.DBHost
	dbPort := c.DBPort
	dbName := c.DBName
	dbUsername := c.DBUsername
	dbPassword := c.DBPassword

	dsn := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s", dbHost, dbPort, dbName, dbUsername, dbPassword)

	db, err = gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		return nil
	}
	return &postgresDB{
		db: db,
	}
}

func (c postgresDB) DB() *gorm.DB {
	return c.db
}
