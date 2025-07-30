package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"github.com/wsb777/call-back/internal/config"
)

type DatabasePG struct {
	DBUsername string
	DBPassword string
	DBName     string
	DBHost     string
	DBPort     string
}

func NewDatabasePG(cfg *config.Config) *DatabasePG {
	return &DatabasePG{
		DBUsername: cfg.DBUser,
		DBPassword: cfg.DBPassword,
		DBName:     cfg.DBName,
		DBHost:     cfg.DBHost,
		DBPort:     cfg.DBPort,
	}
}

func ConnectDBProvider(dbConfig *DatabasePG) (*sql.DB, error) {

	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=%s sslmode=disable",
		dbConfig.DBUsername,
		dbConfig.DBName,
		dbConfig.DBPassword,
		dbConfig.DBHost,
		dbConfig.DBPort,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("ошибка подключения: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("ошибка проверки соединения: %w", err)
	}

	log.Println("Успешное подключение к БД")
	return db, nil
}
