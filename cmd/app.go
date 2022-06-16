package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	log.Println("loading app configuration...")

	// init the config
	config, err := initConfig()
	if err != nil {
		fmt.Print("could not load config: ", err)
		os.Exit(1)
	}

	log.Println("app configuration loaded successfully")
	log.Println("initializing app apiserver...")

	// init the api server
	server, err := initApiServer(config)
	if err != nil {
		fmt.Print("could not init server: ", err)
		os.Exit(1)
	}

	log.Println("app apiserver initialized successfully")
	log.Println("starting app apiserver...")

	// start the server
	err = server.Start()
	if err != nil {
		fmt.Print("could not start server: ", err)
		os.Exit(1)
	}

}
