package web

import (
	"backend/structures"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func Serve(conf map[string]string, db *mongo.Client){

	log.Println("Initialising the web service. Please wait...")

	wa := fiber.New(*&fiber.Config{DisableStartupMessage: true})

	log.Println("Loading modules and setting up the web service.")

	wa.Use(cors.New(cors.Config{
		AllowOrigins: conf["FRONTEND_URL"],
		AllowHeaders:  "Origin, Content-Type, Accept",
	}))

	Route(*&structures.SkyfallenCASApp{ WA: wa, DB: db, Conf:conf })

	log.Println("*********************************************")
	log.Println("Centralized authentication service will be available briefly.")

	err := wa.Listen(":"+conf["WEB_PORT"])

	if err != nil {
		log.Println("An error occurred in during SkyfallenCAS boot.")
		log.Println("Please check the web service settings.")
		log.Println(err)
		log.Fatal("SkyfallenCAS has crashed.")
		return
	}

}