package model

import "time"

type Sesion struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	UsuarioID   uint      `json:"usuario_id"`
	Fecha       time.Time `json:"fecha"`
	DuracionMin int       `json:"duracion_min"`
	Comentarios string    `json:"comentarios"`
}

func (s *Sesion) TableName() string {
	return "sesiones"
}
