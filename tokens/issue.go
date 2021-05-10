package tokens

import (
	"backend/logging"
	"backend/structures"
	"context"
	"github.com/google/uuid"
	"log"
	"time"
)

func Issue(to structures.User, reason string, duration time.Duration, scope string, app structures.SkyfallenCASApp) *structures.AuthToken {

	log.Printf("Received a request from code to issue a new token to %v \n",to.Username)

	tokenStr, err := uuid.NewRandom()

	logging.Fatal(err)

	token := &structures.AuthToken{
		Token:        tokenStr.String(),
		TokenExpires: time.Now().Add(duration),
		TokenCreated: time.Now(),
		TokenScope:   scope,
		TokenReason:  reason,
		TokenUser:    to.Username,
	}

	db := app.DB.Database(app.Conf["DB_CONN_NAME"])

	collection := db.Collection("tokens")

	log.Println("Creating a new context for the Mongo Query.")
	dbctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()


	log.Println("Inserting token into the database.")
	_ , err = collection.InsertOne(dbctx, token)

	logging.Fatal(err)

	log.Println("Successfully issued a token, returning.")

	return token

}
