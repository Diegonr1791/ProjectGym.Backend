package model

import "time"

type SesionEjercicio struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	SesionID     uint      `json:"sesion_id"`
	EjercicioID  uint      `json:"ejercicio_id"`
	Fecha        time.Time `json:"fecha"`
	Series       int       `json:"series"`
	Repeticiones int       `json:"repeticiones"`
	Orden        int       `json:"orden"`
	Peso         float64   `json:"peso"`
	Observacion  string    `json:"observacion"`
}

func (s *SesionEjercicio) TableName() string {
	return "sesion_ejercicios"
}
