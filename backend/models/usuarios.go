package models

import "time"

type Usuarios struct {
	ID              	uint                	`json:"id" gorm:"primary_key"`
	Nombre          	string              	`json:"nombre"`
	Apellido        	string              	`json:"apellido"`
	Dni       			string              	`json:"dni"`
	Telefono        	string              	`json:"telefono"`
	Email       		string              	`json:"email"`
	Fecha_creado      	time.Time             	`json:"fecha_creado"`
	Fecha_actualizado   time.Time               `json:"fecha_actualizado"`
	Estado       		bool              		`json:"estado"`
}

func (Usuarios) TableName() string {
	return "usuarios"
}