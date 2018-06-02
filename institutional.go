package main

import (
	"github.com/paulantezana/institutional/api"
)

func main() {
	//var isMigrate string
	//flag.StringVar(&isMigrate,"migrate","no","migrate models to database [yes/no]")

	// Migrations
	//if isMigrate == "yes" {
	//fmt.Println("Init migration...")
	//migrations.Migrate()
	//fmt.Println("End migration...")
	////}

	// Starting API
	api.Start()
}
