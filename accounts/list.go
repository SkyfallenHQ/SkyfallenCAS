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
	"log"
)

func list(ctx *fiber.Ctx, app structures.SkyfallenCASApp) error {

	logging.Request(ctx)

	twe, twr := tokens.RequestValidation(ctx,app)

	if twe != nil{
		return twr
	}

	db := app.DB.Database(app.Conf["DB_CONN_NAME"])

	collection := db.Collection("users")

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {

		log.Println("An error occurred while serving the /accounts/list endpoint.")
		log.Println(err)

		r := &responses.Basic{Status: "ERR", Error: responses.Error{Code: "500", Description: err.Error()}, Result: "Accounts listing query has failed."}

		resp, err := json.Marshal(r)

		logging.Fatal(err)

		return ctx.Send(resp)

	}

	var users []structures.User

	if err = cursor.All(context.TODO(), &users); err != nil {

		log.Println("An error occurred while serving the /accounts/list endpoint.")
		log.Println(err)

		r := &responses.Basic{Status: "ERR", Error: responses.Error{Code: "500", Description: err.Error()}, Result: "Accounts listing query has failed."}

		resp, err := json.Marshal(r)

		logging.Fatal(err)

		return ctx.Send(resp)

	}

	log.Println("All users successfully fetched.")

	r := &responses.Basic{Status: "OK", Error: responses.Error{Code: "", Description: ""}, Result: users}

	resp, err := json.Marshal(r)

	logging.Fatal(err)

	log.Println("Request satisfied, sending response.")

	return ctx.Send(resp)

}