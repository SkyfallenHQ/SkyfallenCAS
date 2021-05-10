package structures

import (
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

type SkyfallenCASApp struct {

	WA *fiber.App
	DB *mongo.Client
	Conf map[string]string

}
