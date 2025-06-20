package model

import "time"

type Rutina struct {
	ID            uint      `gorm:"primaryKey" json:"id"`
	Nombre        string    `json:"nombre"`
	Objetivo      string    `json:"objetivo"`
	FechaCreacion time.Time `json:"fecha_creacion"`
	Publica       bool      `json:"publica"`
	UsuarioID     uint      `json:"usuario_id"`
}

func (Rutina) TableName() string {
	return "rutinas"
}
