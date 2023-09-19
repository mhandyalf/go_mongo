package main

import (
	"go_mongo/database"
	"go_mongo/routes"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {

	if err := database.InitDB(); err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	e := echo.New()

	routes.SetEmployeeRoutes(e)

	e.Logger.Fatal(e.Start(":8080"))
}
