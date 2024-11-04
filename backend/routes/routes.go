package routes

import (
	rolesController "puchito/controllers/roles"
	generosController "puchito/controllers/generos"
	librosController "puchito/controllers/libros"

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
	b.PUT("/generos/:id", generosController.Set)
	b.DELETE("/generos/:id", generosController.Delete)

	//libros
	b.GET("/libros", librosController.GetAll)
	b.POST("/libros", librosController.Create)

}