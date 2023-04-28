package configurations

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"os"
)

type SqlConnectionConfig struct {
	db               *sql.DB
	connectionString string
}

type SqlConnector interface {
	EstablishConnectionToDatabase() (*sql.DB, error)
}

func EstablishConnectionToDatabase() (*sql.DB, error) {
	connectionString := os.Getenv("MYSQL_ConnectionString")
	fmt.Println("connectionString: ", connectionString)
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		log.Println("problem med connectionString")
		log.Fatal(err)
		return db, errors.New(err.Error())
	}

	if err := db.Ping(); err != nil {
		log.Println("ping database... doesn't look like we are getting in contact")
		log.Fatal(err)
		return db, errors.New(err.Error())
	}

	return db, nil
}
