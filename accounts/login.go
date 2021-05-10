package accounts

import (
	"backend/logging"
	"backend/responses"
	"backend/structures"
	"backend/tokens"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

func login(ctx *fiber.Ctx, app structures.SkyfallenCASApp) error {

	logging.Request(ctx)

	db := app.DB.Database(app.Conf["DB_CONN_NAME"])

	collection := db.Collection("users")

	log.Println("Creating a new context for the Mongo Query.")
	dbctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	var matchingUser structures.User

	err := collection.FindOne(dbctx, bson.M{ "username": ctx.FormValue("username")} ).Decode(&matchingUser)

	if err != nil {

		log.Println("An error occurred while serving the /accounts/login endpoint.")
		log.Println(err)

		r := &responses.Basic{Status: "ERR", Error: responses.Error{Code: "403", Description: err.Error()}, Result: "user_not_exists"}

		resp, err := json.Marshal(r)

		logging.Fatal(err)

		return ctx.Send(resp)

	}

	err = bcrypt.CompareHashAndPassword([]byte(matchingUser.Password), []byte(ctx.FormValue("password")))

	if err == nil {

		log.Println("The login was successful. Now waiting for the token to be issued.")

		token := tokens.Issue(matchingUser,"login", time.Duration(3600 * time.Second),"user", app)

		log.Println("Token created successfully.")

		r := &responses.TokenResponse{Status: "OK", Error: responses.Error{Code: "", Description: ""}, Token: *token}

		resp, err := json.Marshal(r)

		logging.Fatal(err)

		log.Println("Request satisfied, sending response.")

		return ctx.Send(resp)

	} else {

		r := &responses.Basic{Status: "ERR", Error: responses.Error{Code: "403", Description: err.Error()}, Result: "password_is_incorrect"}

		resp, err := json.Marshal(r)

		logging.Fatal(err)

		return ctx.Send(resp)

	}

}
