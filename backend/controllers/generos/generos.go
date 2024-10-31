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

type Genero struct {
	Genero	string	`json:"genero"`
}

func Create(c echo.Context) error {
	db := database.GetDb()
	genero := new(Genero)

	if err := c.Bind(genero); err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "invalid request body " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	if err := db.Raw(`INSERT INTO 'generos' ('genero') values ("?")`, genero).Error; err != nil {
		response := ResponseMessage{
			Status: "error",
			Message: "error creating gender " + err.Error(),
		}
		return c.JSON(http.StatusBadRequest,response)
	}

	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Message: "gender created",
	})
}