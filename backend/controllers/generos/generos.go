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

	//Order By
	if c.QueryParam("sortField") != "" {
		db = db.Order(c.QueryParam("sortField") + " " + c.QueryParam("sortOrder"))
	} else {
		db = db.Order("id")
	}

	var generos []models.Generos

	db.Raw(`SELECT * FROM puchito.generos`).Find(&generos)

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

	var Idgenero uint
	if err := db.Exec(`INSERT INTO puchito.generos (genero) values (?) RETURNING ID`, genero.Genero).Scan(&Idgenero).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error creating gender " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	genero.ID = Idgenero

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