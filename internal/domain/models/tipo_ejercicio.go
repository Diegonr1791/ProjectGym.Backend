package model

type TipoEjercicio struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Nombre string `json:"nombre"`
}

func (TipoEjercicio) TableName() string {
	return "tipo_ejercicio"
}
