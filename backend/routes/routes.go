package routes

import (
	rolesController "puchito/controllers/roles"
	generosController "puchito/controllers/generos"

	jwt "puchito/routes/middleware"

	"github.com/labstack/echo/v4"
)

func InitRoutes(e *echo.Echo) {
	a := e.Group("/api")    //privado
	b := e.Group("/public") //publico

	a.Use(jwt.EchoEnsureValidToken())

	//roles
	b.GET("/roles", rolesController.GetAll)

	//generos
	b.GET("/generos", generosController.GetAll)
	b.POST("/generos", generosController.Create)

}