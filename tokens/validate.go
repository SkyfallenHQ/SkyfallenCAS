package tokens

import (
	"backend/structures"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func Validate(token string, app structures.SkyfallenCASApp) *structures.AuthToken {

	log.Println("Starting validation of a  token.")

	db := app.DB.Database(app.Conf["DB_CONN_NAME"])

	collection := db.Collection("tokens")

	log.Println("Creating a new context for the Mongo Query.")
	dbCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var matchingToken structures.AuthToken

	err := collection.FindOne(dbCtx, bson.M{ "token":  token}).Decode(&matchingToken)

	if err != nil {

		log.Println("The provided token was invalid.")
		return nil

	} else {

		log.Println("A token was found. Checking expiration.")

		if matchingToken.TokenExpires.Before(time.Now()) {

			log.Println("Provided token has expired.")

			return nil

		} else {

			log.Println("The provided token is valid and has not expired.")

			return &matchingToken

		}

	}

}
