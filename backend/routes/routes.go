package routes

import (
	rolesController "puchito/controllers/roles"
	generosController "puchito/controllers/generos"
	librosController "puchito/controllers/libros"
	usuariosController "puchito/controllers/usuarios"
	prestamosController "puchito/controllers/prestamos"

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
	b.GET("/generos/:id", generosController.Get)
	b.PUT("/generos/:id", generosController.Set)
	b.DELETE("/generos/:id", generosController.Delete)

	//libros
	b.GET("/libros", librosController.GetAll)
	b.POST("/libros", librosController.Create)
	b.GET("/libros/:id", librosController.Get)
	b.PUT("/libros/:id", librosController.Set)
	b.DELETE("/libros/:id", librosController.Delete)

	//usuarios
	b.GET("/usuarios", usuariosController.GetAll)
	b.POST("/usuarios", usuariosController.Create)
	b.GET("/usuarios/:id", usuariosController.Get)
	b.PUT("/usuarios/:id", usuariosController.Set)
	b.DELETE("/usuarios/:id", usuariosController.Delete)

	//prestamos
	b.GET("/prestamos", prestamosController.GetAll)
	b.POST("/prestamos", prestamosController.Create)
	b.GET("/prestamos/:id", prestamosController.Get)
	b.PUT("/prestamos/:id", prestamosController.Set)

}