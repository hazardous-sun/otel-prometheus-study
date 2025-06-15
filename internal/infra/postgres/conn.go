package postgres

import (
	"database/sql"
	"fmt"
	"otel-prometheus-study/internal/logger"
)

func ConnectDB() (*sql.DB, error) {
	config := LoadConfig()
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBName)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	logger.LogInfo("Connected to " + config.DBName)
	return db, nil
}
