package config

import "os"

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
	Host     string
	Name     string
	User     string
	Password string
	Port     string
}

func GetDBConfig() DB {
	db := DB{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		Name:     os.Getenv("DB_NAME"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
	}
	return db
}
