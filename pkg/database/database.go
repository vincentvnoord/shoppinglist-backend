package database

import (
	"database/sql"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

var dbConnection *sql.DB

func Connection() (*sql.DB, error) {
	if dbConnection != nil {
		err := dbConnection.Ping()
		if err == nil {
			return dbConnection, nil
		}
	}

	db, err := connectToDatabase()
	if err != nil {
		return nil, err
	}

	dbConnection = db

	return dbConnection, nil
}

func LoadEnv() error {
	knownFile := "go.mod"

	dir, err := os.Getwd()
	if err != nil {
		return err
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, knownFile)); !os.IsNotExist(err) {
			err := godotenv.Load(dir + "/.env")
			if err != nil {
				return err
			}

			return nil
		}

		parentDir := filepath.Dir(dir)
		if parentDir == dir { // Reached the root directory
			break
		}
		dir = parentDir
	}

	return err
}

func connectToDatabase() (*sql.DB, error) {
	dbUser := os.Getenv("DB_USER")
	dbName := os.Getenv("DB_NAME")
	dbPassword := os.Getenv("DB_PASSWORD")

	connStr := "user= " + dbUser + " dbname=" + dbName + " password=" + dbPassword + " sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
