package prestamos

import (
	"net/http"
	"puchito/database"
	"puchito/models"
	"time"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Prestamos []models.Prestamos		`json:"prestamos,omitempty"`
	Prestamo   *models.Prestamos		`json:"prestamo,omitempty"`
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

	var prestamos []models.Prestamos

	db.Raw(`SELECT * FROM puchito.prestamos`).Find(&prestamos)

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
	db.Raw(`SELECT * FROM puchito.prestamos WHERE id = ?`,id).First(&prestamo)

	data := Data{Prestamo: prestamo}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Create(c echo.Context) error {
	//db := database.GetDb()
	prestamo := new(models.Prestamos)

	if err := c.Bind(prestamo); err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	/* if err := db.Exec(`INSERT INTO puchito.prestamos (id_usuario,id_libro,fecha_devolucion_estimada) values (?,?,?)`, libro.Nombre,libro.Autor,libro.Fecha_creado,libro.Id_genero).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error creating gender " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	} */

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Message: "gender created",
	})
}

func Set(c echo.Context) error{
	db := database.GetDb()
	libro := new(models.Libros)

	if err := c.Bind(libro); err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	if err := db.Exec(`UPDATE puchito.libros SET nombre = ? , autor = ? , fecha_lanzamiento = ? , id_genero = ? , fecha_actualizado = ? WHERE id = ?`, libro.Nombre,libro.Autor,libro.Fecha_lanzamiento,libro.Id_genero, time.Now(), c.Param("id")).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error editing book " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Message: "book edited",
	})
}