package model

type RutinaGrupoMuscular struct {
	ID              uint `gorm:"primaryKey" json:"id"`
	RutinaID        uint `json:"rutina_id"`
	GrupoMuscularID uint `json:"grupo_muscular_id"`
}

func (RutinaGrupoMuscular) TableName() string {
	return "rutina_grupo_muscular"
}

type RutinaConGruposMusculares struct {
	ID               uint            `json:"id"`
	RutinaID         uint            `json:"rutina_id"`
	GruposMusculares []GrupoMuscular `json:"grupos_musculares"`
}
