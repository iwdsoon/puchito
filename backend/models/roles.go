package models

type Roles struct {
	ID          uint                	`json:"id" gorm:"primary_key"`
	Rol     	string              	`json:"rol"`
}

func (Roles) TableName() string {
	return "roles"
}