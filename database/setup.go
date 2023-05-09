package database

import (
	"DigiShop/tools"
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

type DBConfig struct {
	DB Config `json:"database"`
}

type Config struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func Setup() *sql.DB {
	config := loadConfig("config/config.json")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database)

	db, err := sql.Open("mysql", dataSourceName)
	defer func(db *sql.DB) {
		err := db.Close()
		tools.CheckError(err, "closing to mysql!")
	}(db)
	tools.CheckError(err, "connecting to mysql!")

	err = db.Ping()
	tools.CheckError(err, "get ping from mysql!")

	fmt.Println("The connection to the database was established.")
	return db
}

func loadConfig(filename string) *DBConfig {
	file, err := os.ReadFile(filename)
	tools.CheckError(err, "open database config!")

	var dbConfig DBConfig
	err = json.Unmarshal(file, &dbConfig)
	tools.CheckError(err, "json unmarshal in load database config!")

	return &dbConfig
}
