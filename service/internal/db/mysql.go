package db

import (
	"database/sql"
	"time"

	"github.com/Seiya-Tagami/pecopeco-service/internal/config"
	"github.com/go-sql-driver/mysql"
)

var sqlDB *sql.DB

func NewMySQL() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
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
		panic(err)
	}

	sqlDB = db
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
