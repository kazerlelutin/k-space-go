package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import du driver PostgreSQL pour CockroachDB
)

func Db() {
	connStr := "postgresql://kazerlelutin:ziKUe3adkXPfrjQHUXjZcg@modern-hisser-2843.jxf.gcp-europe-west1.cockroachlabs.cloud:26257/postgres?sslmode=require"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}

	defer db.Close()

	// Vérification de la connexion
	err = db.Ping()
	if err != nil {
		log.Fatal("Could not connect to CockroachDB: ", err)
	}

	fmt.Println("Connected to CockroachDB successfully!")
}
