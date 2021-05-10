package main

import (
	"backend/config"
	"backend/database"
	"backend/web"
	"fmt"
	"log"
)

func main() {

	fmt.Printf(" ____  _           __       _ _            \n/ ___|| | ___   _ / _| __ _| | | ___ _ __  \n\\___ \\| |/ / | | | |_ / _` | | |/ _ \\ '_ \\ \n ___) |   <| |_| |  _| (_| | | |  __/ | | |\n|____/|_|\\_\\\\__, |_|  \\__,_|_|_|\\___|_| |_|\n            |___/                          \n\n")
	fmt.Println("")

	fmt.Println("**************************************************************")

	//time.Sleep(3 * time.Second)

	log.Println("Welcome to Centralized Authentication Service.")
	log.Println("Service will be booting up shortly.")
	log.Println("(C) 2021 - Skyfallen | All rights reserved.")

	//time.Sleep(2 * time.Second)

	conf := config.Parse()

	log.Printf("Skyfallen CAS %v is starting.", conf["CAS_VERSION"])

	db := database.Connect(conf)

	log.Printf("Starting the web server on port %v", conf["WEB_PORT"])

	web.Serve(conf,db)

}
