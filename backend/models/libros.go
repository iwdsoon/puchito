package models

import "time"

type Libros struct {
	ID              	uint                	`json:"id" gorm:"primary_key"`
	Nombre          	string              	`json:"nombre"`
	Autor       		string              	`json:"autor"`
	Fecha_lanzamiento   time.Time              	`json:"fecha_lanzamiento"`
	Id_genero       	uint              		`json:"id_genero"`
	Fecha_creado      	time.Time             	`json:"fecha_creado"`
	Fecha_actualizado   time.Time               `json:"fecha_actualizado"`
	Estado       		bool              		`json:"estado"`
}

func (Libros) TableName() string {
	return "libros"
}