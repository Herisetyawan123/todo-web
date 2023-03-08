package utils

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"todo-web/entity"
)

var db *gorm.DB

func ConnectDb() error {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DriverName: "pgx",
		DSN:        os.Getenv("DATABASE_URL"),
	}))
	if err != nil {
		panic(err)
	}
	conn.AutoMigrate(entity.User{})
	SetupDBConnection(conn)
	return err
}

func SetupDBConnection(DB *gorm.DB) {
	db = DB
}

func getDBConnection() *gorm.DB {
	return db
}
