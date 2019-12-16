package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
	"testing"
)

func TestReadDataFromPostgres(t *testing.T) {
	pg_host, exists := os.LookupEnv("POSTGRES_HOST")
	if !exists {
		pg_host = "localhost"
	}

	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg_host, 5432, "postgres", "", "")

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to Postgres DB...")

	var codefreshUrl string
	row := db.QueryRow(`SELECT url FROM link WHERE name=$1;`, "Codefresh")

	fmt.Println("Reading Test data...")

	switch err := row.Scan(&codefreshUrl); err {
	case sql.ErrNoRows:
		t.Errorf("No rows were returned!")
	}

	if codefreshUrl != "http://www.codefresh.io" {
		t.Errorf("Got = %s; want http://www.codefresh.io", codefreshUrl)
	}

	fmt.Println("Finished integration test")

}
