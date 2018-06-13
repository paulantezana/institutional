package config

import (
    "github.com/jinzhu/gorm"
    "fmt"
    "log"
)

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