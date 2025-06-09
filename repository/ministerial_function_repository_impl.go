package repository

import (
	"church_ministerial_functions/domain"

	"gorm.io/gorm"
)

type GormMinisterialFunctionRepository struct {
	db *gorm.DB
}

func NewGormMinisterialFunctionRepository(db *gorm.DB) *GormMinisterialFunctionRepository {
	return &GormMinisterialFunctionRepository{db: db}
}

func (r *GormMinisterialFunctionRepository) Save(mf *domain.MinisterialFunction) error {
	return r.db.Create(mf).Error
}

func (r *GormMinisterialFunctionRepository) FindByID(id uint) (*domain.MinisterialFunction, error) {
	var mf domain.MinisterialFunction
	err := r.db.First(&mf, id).Error
	return &mf, err
}

func (r *GormMinisterialFunctionRepository) FindAll() ([]domain.MinisterialFunction, error) {
	var mfs []domain.MinisterialFunction
	err := r.db.Find(&mfs).Error
	return mfs, err
}

func (r *GormMinisterialFunctionRepository) Update(mf *domain.MinisterialFunction) error {
	return r.db.Save(mf).Error
}

func (r *GormMinisterialFunctionRepository) Delete(id uint) error {
	return r.db.Delete(&domain.MinisterialFunction{}, id).Error
}
