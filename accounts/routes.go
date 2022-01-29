package accounts

import (
	"backend/structures"
	"github.com/gofiber/fiber/v2"
)

func Route(app structures.SkyfallenCASApp){

	app.WA.Post("/accounts/create", func(ctx *fiber.Ctx) error {

		return createAccount(ctx, app)

	})

	app.WA.Post("/accounts/login", func(ctx *fiber.Ctx) error {

		return login(ctx, app)

	})


	app.WA.Get("/accounts/list", func(ctx *fiber.Ctx) error {

		return list(ctx, app)

	})

}
