package database

import (
	"database/sql"
	"fmt"
	"log"

	"onlinecourse/internal/config"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDB(cfg config.Config) {
	var err error
	fmt.Println("0")
	// dsn := "host=0.0.0.0 user=webservice_user password=your_strong_password dbname=webservice sslmode=disable"
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.POSTGRESHOST, cfg.POSTGRESUSER, cfg.POSTGRESPASSWORD, cfg.POSTGRESDB, cfg.POSTGRESPORT)

	DB, err = sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatal("Database not responding:", err)
	}

	fmt.Println("Connected to PostgreSQL!")
}
