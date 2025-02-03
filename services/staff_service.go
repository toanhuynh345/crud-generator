package services

import (
    "github.com/toanmars/crud-generator/models"
    "github.com/toanmars/crud-generator/repositories"
)

type StaffService struct {
    repo *repositories.StaffRepository
}

func NewStaffService(repo *repositories.StaffRepository) *StaffService {
    return &StaffService{repo: repo}
}

func (s *StaffService) Create(model *models.Staff) error {
    return s.repo.Create(model)
}

func (s *StaffService) GetByID(id uint) (*models.Staff, error) {
    return s.repo.FindByID(id)
}

func (s *StaffService) Update(model *models.Staff) error {
    return s.repo.Update(model)
}

func (s *StaffService) Delete(id uint) error {
    return s.repo.Delete(id)
}