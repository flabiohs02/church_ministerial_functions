package usecase

import (
	"church_ministerial_functions/domain"
	"church_ministerial_functions/repository"
)

type MinisterialFunctionService struct {
	repo repository.MinisterialFunctionRepository
}

func NewMinisterialFunctionService(repo repository.MinisterialFunctionRepository) *MinisterialFunctionService {
	return &MinisterialFunctionService{repo: repo}
}

func (s *MinisterialFunctionService) CreateMinisterialFunction(mf *domain.MinisterialFunction) error {
	return s.repo.Save(mf)
}

func (s *MinisterialFunctionService) GetMinisterialFunctionByID(id uint) (*domain.MinisterialFunction, error) {
	return s.repo.FindByID(id)
}

func (s *MinisterialFunctionService) GetAllMinisterialFunctions() ([]domain.MinisterialFunction, error) {
	return s.repo.FindAll()
}

func (s *MinisterialFunctionService) UpdateMinisterialFunction(mf *domain.MinisterialFunction) error {
	return s.repo.Update(mf)
}

func (s *MinisterialFunctionService) DeleteMinisterialFunction(id uint) error {
	return s.repo.Delete(id)
}
