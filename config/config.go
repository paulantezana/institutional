package config

import (
	"encoding/json"
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

type Email struct {
    Name string
    From string
    Password string
    Server string
    Host string
}

type Config struct {
	Mode     Mode
	Database Database
	Email Email
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
