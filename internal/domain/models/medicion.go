package model

import "time"

type Medicion struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	UsuarioID     uint      `json:"usuario_id"`
	Fecha         time.Time `json:"fecha"`
	PesoCorporal  float32   `json:"peso_corporal"`
	GrasaCorporal float32   `json:"grasa_corporal"`
	Musculo       float32   `json:"musculo"`
}

func (Medicion) TableName() string {
	return "mediciones"
}
