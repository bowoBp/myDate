package db

import (
	"fmt"
	config "github.com/bowoBp/myDate/pkg/reader"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Default() (*gorm.DB, error) {
	var (
		username string
		password string
		host     string
		port     string
		dbName   string
		sslMode  string
	)

	username = config.GetEnv("PGUSER")
	password = config.GetEnv("PGPASSWORD")
	host = config.GetEnv("PGHOST")
	dbName = config.GetEnv("PGDB")
	port = config.GetEnv("PGPORT")
	sslMode = config.GetEnv("PGSSL")

	dbConn, err := gorm.Open(
		postgres.Open(fmt.Sprintf(
			"user=%s password=%s host=%s port=%s dbname=%s sslmode=%s",
			username,
			password,
			host,
			port,
			dbName,
			sslMode,
		)),
		&gorm.Config{
			CreateBatchSize: 500,
		},
	)
	return dbConn, err
}
