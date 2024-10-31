package models

type Generos struct {
	ID              uint                	`json:"id" gorm:"primary_key"`
	Genero      	string              	`json:"genero"`
}

func (Generos) TableName() string {
	return "generos"
}