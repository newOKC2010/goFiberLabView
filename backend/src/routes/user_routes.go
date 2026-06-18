package routes

import (
	"database/sql"

	manageUser "view_lab/src/controllers/labview/manageUser"
	middleware "view_lab/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app fiber.Router, db *sql.DB) {
	prefix := app.Group("/users",
		middleware.AuthGuards(db, []string{"admin", "super_admin"}),
	)

	prefix.Get("/", manageUser.GetUsers)
	prefix.Post("/update-status", manageUser.UpdateStatus)
}
