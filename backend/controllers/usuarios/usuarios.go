package usuarios

import (
	"net/http"
	"puchito/database"
	"puchito/models"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Usuarios []models.Usuarios		`json:"usuarios,omitempty"`
	Usuario   *models.Usuarios 		`json:"usuario,omitempty"`
}

type ResponseMessage struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAll(c echo.Context) error {
	db := database.GetDb()

	//Order By
	if c.QueryParam("sortField") != "" {
		db = db.Order(c.QueryParam("sortField") + " " + c.QueryParam("sortOrder"))
	} else {
		db = db.Order("id")
	}

	var usuarios []models.Usuarios

	db.Exec(`SELECT * FROM puchito.usuarios WHERE estado = true`).Find(&usuarios)

	data := Data{Usuarios: usuarios}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Get(c echo.Context) error {
	db := database.GetDb()
	id := c.Param("id")

	usuario := new(models.Usuarios)
	db.Exec(`SELECT * FROM puchito.usuarios WHERE id = ?`,id).First(&usuario)

	data := Data{Usuario: usuario}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Create(c echo.Context) error {
	db := database.GetDb()
	usuario := new(models.Usuarios)

	if err := c.Bind(usuario); err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	if err := db.Exec(`INSERT INTO puchito.usuarios (nombre,apellido,dni,telefono,email,id_rol) values (?,?,?,?,?,?)`, usuario.Nombre,usuario.Apellido,usuario.Dni,usuario.Telefono,usuario.Email,2).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error creating user " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Message: "user created",
	})
}

func Set(c echo.Context) error{
	db := database.GetDb()
	usuario := new(models.Usuarios)

	if err := c.Bind(usuario); err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	if err := db.Exec(`UPDATE puchito.usuarios SET nombre = ? , apellido = ? , dni = ? , telefono = ? , email = ? WHERE id = ?`, usuario.Nombre,usuario.Apellido,usuario.Dni,usuario.Telefono,usuario.Email, c.Param("id")).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error editing user " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Message: "user edited",
	})
}

func Delete(c echo.Context) error {
	db := database.GetDb()

	if err := db.Exec(`UPDATE puchito.usuario SET estado = false WHERE id = ?`, c.Param("id")).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error deleting" + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}
	
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Message: "user deleted",
	})
} 