package main

import (
	"puchito/config"
	"puchito/database"
	routes "puchito/routes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Echo instance
	e := echo.New()
	e.Use(middleware.Static("/static"))
	e.Static("/api/static", "static")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))

	config.LoadEnvProps("env.properties")
	database.InitDb(config.GetString("dbStr"))

	//API routes
	routes.InitRoutes(e)

	// Start server
	e.Logger.Fatal(e.Start(config.GetString("httpPort")))
}