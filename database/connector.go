package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func Connect(conf map[string]string) *mongo.Client {

	log.Println("Connecting to the database.")

	if conf["DB_TYPE"] != "mongo" {

		log.Println("An error has occurred in the SkyfallenCAS boot.")
		log.Println("SkyfallenCAS only supports MongoDB at the moment.")
		log.Fatal("SkyfallenCAS has crashed.")

	}

	log.Println("Database connection timeout is set to 10 seconds")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Printf("Trying to connect to the %v database server. \n", conf["DB_TYPE"])
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(conf["DB_CONN_URI"]))

	if err != nil {
		log.Println("An error occurred in during SkyfallenCAS boot.")
		log.Println("Please check the database configuration")
		log.Println(err)
		log.Fatal("SkyfallenCAS has crashed.")
		return nil
	}

	log.Printf("SkyfallenCAS has successfully connected to the %v database. \n", conf["DB_TYPE"])

	return client

}
