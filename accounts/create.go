package accounts

import (
	"backend/logging"
	"backend/responses"
	"backend/structures"
	"context"
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"time"
)

func createAccount(ctx *fiber.Ctx, app structures.SkyfallenCASApp) error{

	logging.Request(ctx)

	db := app.DB.Database(app.Conf["DB_CONN_NAME"])

	collection := db.Collection("users")

	u := new(structures.User)

	err := ctx.BodyParser(u)

	if err != nil {

		log.Println("An error occurred while serving the /accounts/create endpoint.")
		log.Println(err)

		r := &responses.Basic{Status: "ERR", Error: responses.Error{Code: "500", Description: err.Error()}, Result: "The user creation has failed."}

		resp, err := json.Marshal(r)

		logging.Fatal(err)

		return ctx.Send(resp)

	}

	err = u.Validate()
	if err != nil {

		log.Println("An error occurred while serving the /accounts/create endpoint.")
		log.Println(err)

		r := &responses.Basic{Status: "ERR", Error: responses.Error{Code: "500", Description: err.Error()}, Result: "The user creation has failed."}

		resp, err := json.Marshal(r)

		logging.Fatal(err)

		return ctx.Send(resp)

	}

	r := &responses.Basic{Status: "OK", Error: responses.Error{Code: "", Description: ""}, Result: "The user has been successfully created."}

	resp, err := json.Marshal(r)

	logging.Fatal(err)
	log.Println("Creating a new context for the Mongo Query.")
	dbctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	log.Println("New account query has successfully passed input validation, checking data conflicts.")
	var matchingUser structures.User

	err = collection.FindOne(dbctx, bson.M{ "username": ctx.FormValue("username")} ).Decode(&matchingUser)

	if err == nil {

		log.Println("An error occurred while serving the /accounts/create endpoint.")
		log.Println("The username already exists.")

		r := &responses.Basic{Status: "ERR", Error: responses.Error{Code: "500", Description: "The username is already used."}, Result: "username_exists"}

		resp, err := json.Marshal(r)

		logging.Fatal(err)

		return ctx.Send(resp)

	}

	err = collection.FindOne(dbctx, bson.M{ "email": ctx.FormValue("email")} ).Decode(&matchingUser)

	if err == nil {

		log.Println("An error occurred while serving the /accounts/create endpoint.")
		log.Println("The email already exists.")

		r := &responses.Basic{Status: "ERR", Error: responses.Error{Code: "500", Description: "The email is already used."}, Result: "email_exists"}

		resp, err := json.Marshal(r)

		logging.Fatal(err)

		return ctx.Send(resp)

	}

	_, err = collection.InsertOne(dbctx, u)
	if err != nil {

		log.Println("An error occurred while serving the /accounts/create endpoint.")
		log.Println(err)

		r := &responses.Basic{Status: "ERR", Error: responses.Error{Code: "500", Description: err.Error()}, Result: "The user creation has failed."}

		resp, err := json.Marshal(r)

		logging.Fatal(err)

		return ctx.Send(resp)

	}

	log.Printf("Request to create a user has succeeded.")

	return ctx.Send(resp)

}
