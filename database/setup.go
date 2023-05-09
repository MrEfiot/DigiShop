package database

import (
	"DigiShop/tools"
	"encoding/json"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

func SetupDB() *gorm.DB {
	config := loadConfig(tools.ConfigAddress)

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.DB.Username, config.DB.Password, config.DB.Host, config.DB.Port, config.DB.Database)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	tools.CheckError(err, "connecting to database!")

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
