package models

import "time"

type Prestamos struct {
	ID              			uint                `json:"id" gorm:"primary_key"`
	Id_usuario          		uint             	`json:"id_usuario"`
	Id_libro       				uint             	`json:"id_libro"`
	Fecha_prestamo     			time.Time           `json:"fecha_prestamo"`
	Fecha_devolucion_estimada   time.Time           `json:"fecha_devolucion_estimada"`
	Fecha_devolucion       		time.Time           `json:"fecha_devolucion"`
}

func (Prestamos) TableName() string {
	return "prestamos"
}