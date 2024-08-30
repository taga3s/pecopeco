package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/taga3s/pecopeco-service/internal/config"
)

var sqlDB *sql.DB

func NewMySQL() error {
	dsn, err := config.GetDSN()
	if err != nil {
		return err
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		err := fmt.Errorf("Error occurred while loading the 'Asia/Tokyo' time zone: %v", err)
		return err
	}

	sqlDB = db

	return nil
}

func GetDB() *sql.DB {
	return sqlDB
}

func CloseDB() {
	sqlDB.Close()
}

func CheckConnection() error {
	if err := sqlDB.Ping(); err != nil {
		return err
	}
	return nil
}
