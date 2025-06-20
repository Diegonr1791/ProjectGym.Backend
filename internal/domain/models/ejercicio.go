package model

type Ejercicio struct {
	ID              uint   `gorm:"primaryKey" json:"id"`
	Nombre          string `json:"nombre"`
	TipoEjercicioID uint   `json:"tipo_ejercicio_id"`
	GrupoMuscularID uint   `json:"grupo_muscular_id"`
}

func (e *Ejercicio) TableName() string {
	return "ejercicios"
}
