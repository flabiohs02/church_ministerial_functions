package repository

import (
	"church_ministerial_functions/domain"
)

type MinisterialFunctionRepository interface {
	Save(mf *domain.MinisterialFunction) error
	FindByID(id uint) (*domain.MinisterialFunction, error)
	FindAll() ([]domain.MinisterialFunction, error)
	Update(mf *domain.MinisterialFunction) error
	Delete(id uint) error
}
