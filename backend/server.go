package main

import (
	"log"

	cors "view_lab/src/controllers/cors"
	conn "view_lab/src/database/connection"
	loadenv "view_lab/src/loadenv"
	routes "view_lab/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	port := loadenv.LoadPort()
	if port == "" {
		log.Fatal("PORT is not set")
	}

	conn.ConnectDB()

	app := fiber.New()
	cors.CorsConfig(app)

	log.Printf("Server started on port: %s", port)

	routes.SetupAuthRoutes(app, conn.DB)
	// routes.SetupEcobaseRoutes(app, conn.Db)
	// routes.SetupDashboardRoutes(app, conn.Db)

	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
