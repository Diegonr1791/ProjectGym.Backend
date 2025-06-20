package model

type GrupoMuscular struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Nombre string `json:"nombre"`
}

func (GrupoMuscular) TableName() string {
	return "grupos_musculares"
}
