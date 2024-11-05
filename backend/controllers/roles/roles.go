package roles

import (
	"puchito/database"
	"puchito/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Data struct {
	Roles []models.Roles	`json:"roles,omitempty"`
	Rol   *models.Roles 	`json:"rol,omitempty"`
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

	var roles []models.Roles

	db.Exec(`SELECT * FROM puchito.roles`).Find(&roles)

	data := Data{Roles: roles}
	return c.JSON(http.StatusOK, ResponseMessage{
		Status: "success",
		Data:   data,
	})
}

/* func GetRolId(Rol string) uint {
	db := database.GetDb()
	rol := new(models.Roles)

	db.Where("rol = ?", Rol).First(&rol)
	return rol.ID
}

func GetRol(Id uint) string {
	db := database.GetDb()
	rol := new(models.Roles)

	db.Where("id = ?", Id).First(&rol)
	return rol.Rol
} */