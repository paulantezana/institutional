package config

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"log"
	"os"
)

// Config models
type Database struct {
	Server   string
	Port     string
	User     string
	Pass     string
	Database string
}

type Mode struct {
	Dev bool
}

type Config struct {
	Mode     Mode
	Database Database
}

// GetConfig return configuration from database json
func GetConfig() Config {
	var c Config

	file, err := os.Open("./config/config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatal(err)
	}

	return c
}

// GetConnection get connection database
func GetConnection() *gorm.DB {
	c := GetConfig()

	// Connection database
	var dsn string
	if c.Mode.Dev {
		// Mode dev in local machine
		dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable", c.Database.User, c.Database.Pass, c.Database.Server, c.Database.Port, c.Database.Database)
	} else {
		// Mode Production in HEROKu
		dsn = fmt.Sprintf("postgres://drvzldhzlyrtra:8a6a02fa8d55a65d5a370330d51b211aef066343e37d968eeb6a46527ba6ee48@ec2-54-225-107-174.compute-1.amazonaws.com:5432/d3mp2040aqlpme")
	}

	//dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.User, c.Pass, c.Server, c.Port, c.Database)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
