package routes

import (
	provider "view_lab/src/controllers/auth/login/provider"

	"github.com/gofiber/fiber/v2"
)

func SetupProviderRoutes(app fiber.Router) {
	prefix := app.Group("/provider")

	prefix.Get("/login", provider.ProviderLogin)
	prefix.Get("/callback", provider.ProviderCallback)
}
