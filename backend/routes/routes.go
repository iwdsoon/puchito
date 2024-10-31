package routes

import (
	rolesController "puchito/controllers/roles"

	jwt "puchito/routes/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	a := e.Group("/api")    //privado
	b := e.Group("/public") //publico

	a.Use(jwt.EchoEnsureValidToken())

	//eventos
	b.GET("/roles", rolesController.GetAll)

}