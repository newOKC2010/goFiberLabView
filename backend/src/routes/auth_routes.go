package routes

import (
	"database/sql"
	"time"

	loginMain "view_lab/src/controllers/auth/login/otp"
	provider "view_lab/src/controllers/auth/login/provider"
	middleware "view_lab/src/middleware"
	ratelimit "view_lab/src/middleware/rateLimit"

	"github.com/gofiber/fiber/v2"
)

func SetupProviderRoutes(app fiber.Router) {
	prefix := app.Group("/provider")

	prefix.Get("/login", provider.ProviderLogin)
	prefix.Get("/callback", provider.ProviderCallback)
}

func SetupAuthRoutes(app fiber.Router, db *sql.DB) {
	prefix := app.Group("/auth")

	prefix.Post("/req", ratelimit.RateLimitByEmail(10, 10*time.Minute), loginMain.RequestOTP(db))
	prefix.Post("/verify", ratelimit.RateLimitByEmail(10, 10*time.Minute), loginMain.VerifyOTP(db))
	prefix.Get("/status", middleware.AuthGuards(db, nil), func(c *fiber.Ctx) error {
		user := c.Locals("user_view_lab").(*middleware.UserViewLabInfo)
		return c.JSON(struct {
			Success bool                        `json:"success"`
			User    *middleware.UserViewLabInfo `json:"user"`
		}{
			Success: true,
			User:    user,
		})
	})

}
