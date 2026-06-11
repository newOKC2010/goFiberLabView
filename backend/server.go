package main

import (
	"log"

	conn "view_lab/src/database/connection"
	loadenv "view_lab/src/loadenv"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := loadenv.LoadPort()
	if port == "" {
		log.Fatal("PORT is not set")
	}

	conn.ConnectDB()

	app := fiber.New()
	// cors.CorsConfig(app)

	log.Printf("Server started on port: %s", port)

	// routes.SetupAuthRoutes(app, conn.Db)
	// routes.SetupEcobaseRoutes(app, conn.Db)
	// routes.SetupDashboardRoutes(app, conn.Db)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
