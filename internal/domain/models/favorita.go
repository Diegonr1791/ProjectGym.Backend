package models

import "time"

type Favorita struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	UsuarioID uint      `json:"usuario_id"`
	RutinaID  uint      `json:"rutina_id"`
	Fecha     time.Time `json:"fecha"`
}

func (Favorita) TableName() string {
	return "favoritas"
}
