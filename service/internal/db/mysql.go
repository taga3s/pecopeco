package db

import (
	"database/sql"
	"time"

	"github.com/Seiya-Tagami/pecopeco-service/internal/config"
	"github.com/go-sql-driver/mysql"
)

func NewMySQL() (*sql.DB, error) {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
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
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
