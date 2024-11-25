package usuarios

import (
	"net/http"
	"puchito/database"
	"puchito/models"
	"time"

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
	
	var usuarios []models.Usuarios

	db.Raw(`SELECT * FROM puchito.usuarios ORDER BY id`).Find(&usuarios)

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
	db.Raw(`SELECT * FROM puchito.usuarios WHERE id = ?`,id).First(&usuario)

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

	if err := db.Exec(`INSERT INTO puchito.usuarios (nombre,apellido,dni,telefono,email,id_rol) values (?,?,?,?,?,?)`, usuario.Nombre,usuario.Apellido,usuario.Dni,usuario.Telefono,usuario.Email,usuario.Id_rol).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error creating user " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	var Idusuario uint
	db.Raw(`SELECT LAST_INSERT_ID()`).First(&Idusuario)
	db.Raw(`SELECT * FROM puchito.usuarios WHERE id = ?`,Idusuario).First(&usuario)
	
	data := Data{Usuario: usuario}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data: data,
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

	if err := db.Exec(`UPDATE puchito.usuarios SET nombre = ? , apellido = ? , dni = ? , telefono = ? , email = ?, fecha_actualizado = ? WHERE id = ?`, usuario.Nombre,usuario.Apellido,usuario.Dni,usuario.Telefono,usuario.Email, time.Now(), c.Param("id")).Error; err != nil {
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

	if err := db.Exec(`UPDATE puchito.usuarios SET estado = false WHERE id = ?`, c.Param("id")).Error; err != nil {
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