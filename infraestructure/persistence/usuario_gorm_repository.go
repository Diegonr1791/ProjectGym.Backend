package persistence

import (
	model "github.com/Diegonr1791/GymBro/internal/domain/models"
	"gorm.io/gorm"
)

type usuarioGormRepository struct {
	DB *gorm.DB
}

func NewUsuarioGormRepository(db *gorm.DB) *usuarioGormRepository {
	return &usuarioGormRepository{DB: db}
}

func (r *usuarioGormRepository) GetAll() ([]model.Usuario, error) {
	var usuarios []model.Usuario
	if err := r.DB.Find(&usuarios).Error; err != nil {
		return nil, err
	}
	return usuarios, nil
}

func (r *usuarioGormRepository) GetByID(id uint) (*model.Usuario, error) {
	var usuario model.Usuario
	if err := r.DB.First(&usuario, id).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *usuarioGormRepository) GetByEmail(email string) (*model.Usuario, error) {
	var usuario model.Usuario
	if err := r.DB.Where("email = ?", email).First(&usuario).Error; err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *usuarioGormRepository) Create(usuario *model.Usuario) error {
	return r.DB.Create(usuario).Error
}

func (r *usuarioGormRepository) Update(usuario *model.Usuario) error {
	return r.DB.Save(usuario).Error
}

func (r *usuarioGormRepository) Delete(id uint) error {
	return r.DB.Delete(&model.Usuario{}, id).Error
}
