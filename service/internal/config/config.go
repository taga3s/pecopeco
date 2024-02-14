package config

import (
	"fmt"
	"os"
	"time"

	"github.com/go-sql-driver/mysql"
)

type Server struct {
	Port string
}

func GetServerConfig() Server {
	server := Server{
		Port: os.Getenv("PORT"),
	}
	return server
}

type DB struct {
	ENV      string
	Host     string
	Name     string
	User     string
	Password string
	Port     string
	DSN_PD   string
}

func GetDSN() (string, error) {
	conf := getDBConfig()

	// 本番用
	if conf.ENV == "pd" {
		return conf.DSN_PD, nil
	}

	// 開発用
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		err := fmt.Errorf("Error occurred while loading the 'Asia/Tokyo' time zone: %v", err)
		return "", err
	}
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
	return c.FormatDSN(), nil
}

func getDBConfig() DB {
	db := DB{
		ENV:      os.Getenv("GO_ENV"),
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		DSN_PD:   os.Getenv("DB_DSN_PD"),
	}
	return db
}

type OIDC struct {
	ClientID string
}

func GetOIDC() OIDC {
	oidc := OIDC{
		ClientID: os.Getenv("OAUTH_CLIENT_ID"),
	}
	return oidc
}
