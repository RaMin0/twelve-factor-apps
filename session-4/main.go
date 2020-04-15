package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Config struct {
	Driver   string
	Host     string
	Port     int
	Username string
	Password string
	Database string
}

func main() {
	var c Config

	loadConfigFromEnv(&c)

	db, err := sql.Open(c.Driver, fmt.Sprintf("%s://%s:%s@%s:%d/%s?sslmode=disable",
		c.Driver, c.Username, c.Password, c.Host, c.Port, c.Database))
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query(`SELECT id, name FROM users`)
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var (
			id   int
			name string
		)
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}

		log.Printf("[%d] %s", id, name)
	}
}

func loadConfigFromEnv(c *Config) {
	c.Driver = os.Getenv("DB_DRIVER")
	c.Host = os.Getenv("DB_HOST")
	c.Port, _ = strconv.Atoi(os.Getenv("DB_PORT"))
	c.Username = os.Getenv("DB_USERNAME")
	c.Password = os.Getenv("DB_PASSWORD")
	c.Database = os.Getenv("DB_DATABASE")
}
