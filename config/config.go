package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	// mysql imported as a driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

// Config ->project config
type Config struct {
	Server   string
	Port     string
	User     string
	Password string
	DataBase string
}

func getConfig() Config {
	var c Config

	file, err := os.Open("./config.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// GetConn -> database connection
func GetConn() *gorm.DB {
	c := getConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", c.User, c.Password, c.Server, c.Port, c.DataBase)

	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
