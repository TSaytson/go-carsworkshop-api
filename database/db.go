package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewDatabase() *sql.DB {
	var envs map[string]string

	envs, err := godotenv.Read(".env")

	if err != nil {
		log.Fatalf("Error loading .env")
	}

	ConnectionString := envs["DATABASE_URL"]

	db, err := sql.Open("mysql", ConnectionString)

	if err != nil {
		panic(err.Error())
	}

	_, err = db.Exec("CREATE DATABASE IF NOT EXISTS carsworkshop;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("use carsworkshop;")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS `cars` (`id` INT AUTO_INCREMENT PRIMARY KEY NOT NULL, `manufacturer` VARCHAR(15) NOT NULL, `model` VARCHAR(50) NOT NULL, `year` INT NOT NULL);")
	if err != nil {
		panic(err)
	}

	return db
}
