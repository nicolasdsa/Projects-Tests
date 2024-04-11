package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
		 const (
			host     = "localhost"
			port     = 5432
			user     = "root"
			password = "password"
			dbname   = "root"
	)

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
        host, port, user, password, dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
			log.Fatal(err)
	}
	//defer db.Close()

	err = db.Ping()
	if err != nil {
			log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database!")
	return db
}