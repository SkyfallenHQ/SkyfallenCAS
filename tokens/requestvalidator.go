package tokens

import (
	"backend/logging"
	"backend/responses"
	"backend/structures"
	"encoding/json"
	"errors"
	"github.com/gofiber/fiber/v2"
	"log"
)

func RequestValidation(ctx *fiber.Ctx, app structures.SkyfallenCASApp) (error, error) {

	log.Println("Request received to validate token.")
	if Validate(ctx.Get("Bearer"), app) == nil{
		log.Println("An error occurred while validating the token.")

		r := &responses.Basic{Status: "ERR", Error: responses.Error{Code: "403", Description: "Authentication has failed."}, Result: "Your query was not processed."}

		resp, err := json.Marshal(r)

		logging.Fatal(err)

		return errors.New("authentication has failed"), ctx.Send(resp)
	}
	return nil, nil

}
