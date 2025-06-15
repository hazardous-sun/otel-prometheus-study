package postgres

import (
	"os"
	"otel-prometheus-study/internal/logger"
	"strconv"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	DBName   string
}

func LoadConfig() Config {
	defaultDBHost := "localhost"
	defaultDBPort := 5432
	defaultDBUser := "postgres"
	defaultDBPassword := "1234"
	defaultDBName := "postgres"

	Host := os.Getenv("DB_HOST")
	if Host == "" {
		logger.LogDebug("DB_HOST environment variable not set, cascading to default", "default", defaultDBHost)
		Host = defaultDBHost
	}

	var Port int
	inputPort := os.Getenv("DB_PORT")
	if inputPort == "" {
		logger.LogDebug("DB_PORT environment variable not set, cascading to default", "default", defaultDBPort)
		Port = defaultDBPort
	} else {
		var err error
		Port, err = strconv.Atoi(inputPort)
		if err != nil {
			logger.LogDebug("DB_PORT environment variable not set, cascading to default", "default", inputPort)
			Port = defaultDBPort
		}
	}

	User := os.Getenv("DB_USER")
	if User == "" {
		logger.LogDebug("DB_USER environment variable not set, cascading to default", "default", defaultDBUser)
		User = defaultDBUser
	}

	Password := os.Getenv("DB_PASSWORD")
	if Password == "" {
		logger.LogDebug("DB_PASSWORD environment variable not set, cascading to default", "default", defaultDBPassword)
		Password = defaultDBPassword
	}

	DBName := os.Getenv("DB_NAME")
	if DBName == "" {
		logger.LogDebug("DB_NAME environment variable not set, cascading to default", "default", defaultDBName)
		DBName = defaultDBName
	}

	config := Config{
		Host:     Host,
		Port:     Port,
		User:     User,
		Password: Password,
		DBName:   DBName,
	}
	return config
}
