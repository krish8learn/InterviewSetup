package model

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// EnvKeys holds all environment configuration values
type EnvKeys struct {
	MongoURI string
	MySQLDSN string
}

// LoadEnv loads environment variables and returns an EnvKeys instance
func LoadEnv() *EnvKeys {
	if err := godotenv.Load(); err != nil {
		log.Println("no .env file found, reading from environment")
	}

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		log.Fatal("MONGO_URI is not set")
	}

	mysqlDSN := os.Getenv("MYSQL_DSN")
	if mysqlDSN == "" {
		mysqlDSN = "root:secret@tcp(localhost:3306)/?parseTime=true"
		log.Println("MYSQL_DSN not set, using default:", mysqlDSN)
	}

	return &EnvKeys{
		MongoURI: mongoURI,
		MySQLDSN: mysqlDSN,
	}
}
