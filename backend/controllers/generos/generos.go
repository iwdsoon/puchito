package generos

import (
	"net/http"
	"puchito/database"
	"puchito/models"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Generos []models.Generos	`json:"generos,omitempty"`
	Genero   *models.Generos 	`json:"genero,omitempty"`
}

type ResponseMessage struct {
	Status  string `json:"status"`
	Data    any    `json:"data,omitempty"`
	Message string `json:"message,omitempty"`
}

func GetAll(c echo.Context) error {
	db := database.GetDb()
	
	var generos []models.Generos

	db.Raw(`SELECT * FROM puchito.generos ORDER BY id`).Find(&generos)

	data := Data{Generos: generos}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

func Get(c echo.Context) error {
	db := database.GetDb()
	id := c.Param("id")

	genero := new(models.Generos)
	db.Raw(`SELECT * FROM puchito.generos WHERE id = ?`,id).First(&genero)

	data := Data{Genero: genero}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

type Genero struct {
	Gender	string	`json:"genero"`
}

func Create(c echo.Context) error {
	db := database.GetDb()
	genero := new(models.Generos)

	if err := c.Bind(genero); err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	if err := db.Exec(`INSERT INTO puchito.generos (genero) values (?)`, genero.Genero).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error creating gender " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	var Idgenero uint
	db.Raw(`SELECT LAST_INSERT_ID()`).First(&Idgenero)
	db.Raw(`SELECT * FROM puchito.generos WHERE id = ?`,Idgenero).First(&genero)

	data := Data{Genero: genero}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data: data,
		Message: "gender created",
	})
}

func Set(c echo.Context) error{
	db := database.GetDb()
	genero := new(models.Generos)

	if err := c.Bind(genero); err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	if err := db.Exec(`UPDATE puchito.generos SET genero = ? WHERE id = ?`, genero.Genero, c.Param("id")).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error editing gender " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Message: "gender edited",
	})
}

func Delete(c echo.Context) error {
	db := database.GetDb()

	if err := db.Exec(`DELETE FROM puchito.generos WHERE id = ?`, c.Param("id")).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error deleting" + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}
	
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Message: "gender deleted",
	})
} 