package prestamos

import (
	"net/http"
	"puchito/database"
	"puchito/models"
	"time"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Prestamos []models.Prestamos `json:"prestamos,omitempty"`
	Prestamo  *models.Prestamos  `json:"prestamo,omitempty"`
}

type ResponseMessage struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAll(c echo.Context) error {
	db := database.GetDb()

	var prestamos []models.Prestamos

	db.Raw(`SELECT * FROM puchito.prestamos ORDER BY id`).Find(&prestamos)

	data := Data{Prestamos: prestamos}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Get(c echo.Context) error {
	db := database.GetDb()
	id := c.Param("id")

	prestamo := new(models.Prestamos)
	db.Raw(`SELECT * FROM puchito.prestamos WHERE id = ?`, id).First(&prestamo)

	data := Data{Prestamo: prestamo}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Create(c echo.Context) error {
	db := database.GetDb()
	prestamo := new(models.Prestamos)

	if err := c.Bind(prestamo); err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	var usuario models.Usuarios
	if err := db.Raw(`SELECT * FROM puchito.usuarios WHERE id = ? AND estado = 1`, prestamo.Id_usuario).First(&usuario).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "user not found or inactive",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	var libro models.Libros
	if err := db.Raw(`SELECT * FROM puchito.libros WHERE id = ? AND estado = 1`, prestamo.Id_libro).First(&libro).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "book not found or inactive",
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Exec(`INSERT INTO puchito.prestamos (id_usuario,id_libro,fecha_devolucion_estimada) values (?,?,?)`, prestamo.Id_usuario, prestamo.Id_libro, time.Now().Add((7 * 24 * time.Hour))).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "error creating prestamo " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	var Idprestamo uint
	db.Raw(`SELECT LAST_INSERT_ID()`).First(&Idprestamo)
	db.Raw(`SELECT * FROM puchito.prestamos WHERE id = ?`, Idprestamo).First(&prestamo)

	data := Data{Prestamo: prestamo}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status:  "success",
		Data:    data,
		Message: "prestamo created",
	})
}

func Set(c echo.Context) error {
	db := database.GetDb()

	if err := db.Exec(`UPDATE puchito.usuarios u 
	JOIN puchito.prestamos p ON p.id_usuario = u.id
	SET u.estado = CASE WHEN p.fecha_devolucion_estimada < NOW() THEN 0 ELSE u.estado END
	WHERE p.id = ?`, c.Param("id")).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "error updating user estado: " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	if err := db.Exec(`UPDATE puchito.prestamos SET fecha_devolucion = ? WHERE id = ?`, time.Now(), c.Param("id")).Error; err != nil {
		response := ResponseMessage{
			Status:  "error",
			Message: "error updating prestamo " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest, response)
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status:  "success",
		Message: "prestamo updated",
	})
}
