package logging

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

func Request(ctx *fiber.Ctx) {

	log.Println("Web API request received.")
	log.Println(ctx.IP() + " has made a request to "+string(ctx.Request().RequestURI()))

}