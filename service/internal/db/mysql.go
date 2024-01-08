package db

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/Seiya-Tagami/pecopeco-service/internal/config"
	"github.com/go-sql-driver/mysql"
)

var sqlDB *sql.DB

func NewMySQL() error {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		err := fmt.Errorf("Error occurred while loading the 'Asia/Tokyo' time zone: %v", err)
		return err
	}
	conf := config.GetDBConfig()
	c := mysql.Config{
		DBName:    conf.Name,
		User:      conf.User,
		Passwd:    conf.Password,
		Addr:      conf.Host + ":" + conf.Port,
		Net:       "tcp",
		ParseTime: true,
		Collation: "utf8mb4_0900_ai_ci",
		Loc:       jst,
	}
	db, err := sql.Open("mysql", c.FormatDSN())
	if err != nil {
		err := fmt.Errorf("Error occurred while loading the 'Asia/Tokyo' time zone: %v", err)
		return err
	}

	sqlDB = db

	return nil
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
