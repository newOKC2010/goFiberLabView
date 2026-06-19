package routes

import (
	"database/sql"

	resultLab "view_lab/src/controllers/labview/resultLab"
	middleware "view_lab/src/middleware"

	"github.com/gofiber/fiber/v2"
)

func SetupLabRoutes(app fiber.Router, db *sql.DB) {
	prefix := app.Group("/lab",
		middleware.AuthGuards(db, []string{"user", "admin", "super_admin"}),
	)

	prefix.Post("/results", resultLab.GetLabResults)
}
