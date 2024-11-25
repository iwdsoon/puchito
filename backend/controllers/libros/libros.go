package libros

import (
	"net/http"
	"puchito/database"
	"puchito/models"
	"time"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Libros []models.Libros		`json:"libros,omitempty"`
	Libro   *models.Libros 		`json:"libro,omitempty"`
}

type ResponseMessage struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}


func GetAll(c echo.Context) error {
	db := database.GetDb()
	
	var libros []models.Libros

	db.Raw(`SELECT * FROM puchito.libros ORDER BY id`).Find(&libros)

	data := Data{Libros: libros}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Get(c echo.Context) error {
	db := database.GetDb()
	id := c.Param("id")

	libro := new(models.Libros)
	db.Raw(`SELECT * FROM puchito.libros WHERE id = ?`,id).First(&libro)

	data := Data{Libro: libro}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Create(c echo.Context) error {
	db := database.GetDb()
	libro := new(models.Libros)

	if err := c.Bind(libro); err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	if err := db.Exec(`INSERT INTO puchito.libros (nombre,autor,fecha_lanzamiento,id_genero) values (?,?,?,?)`, libro.Nombre,libro.Autor,libro.Fecha_lanzamiento,libro.Id_genero).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error creating book " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	var Idlibro uint
	db.Raw(`SELECT LAST_INSERT_ID()`).First(&Idlibro)
	db.Raw(`SELECT * FROM puchito.libros WHERE id = ?`,Idlibro).First(&libro)
	
	data := Data{Libro: libro}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data: data,
		Message: "book created",
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

func Delete(c echo.Context) error {
	db := database.GetDb()

	if err := db.Exec(`UPDATE puchito.libros SET estado = false WHERE id = ?`, c.Param("id")).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error deleting" + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}
	
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Message: "book deleted",
	})
} 