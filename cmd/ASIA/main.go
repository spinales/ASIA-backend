package main

import (
	"log"

	"github.com/spinales/ASIA-backend/handlers"
	"github.com/spinales/ASIA-backend/util"
)

func main() {
	config := util.LoadConfig()

	db, err := util.InitDB(&config)
	if err != nil {
		log.Fatal("failed to connect database\n", err)
	}

	err = util.MigrateDB("/db/migration", &config)
	if err != nil {
		log.Fatal("Error executing migration\n", err)
	}

	server, err := handlers.NewServer(db, &config)
	if err != nil {
		log.Fatalln("Cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatalln("Cannot start server: ", err)
	}
}
