package models

type Usuario struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	Nombre       string         `json:"nombre"`
	Email        string         `gorm:"unique" json:"email"`
	Password     string         `json:"-"`
	RefreshToken []RefreshToken `gorm:"foreignKey:UserID" json:"-"`
}
