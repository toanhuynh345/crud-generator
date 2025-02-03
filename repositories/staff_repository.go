package repositories

import (
    "github.com/toanmars/crud-generator/models"
    "gorm.io/gorm"
)

type StaffRepository struct {
    db *gorm.DB
}

func NewStaffRepository(db *gorm.DB) *StaffRepository {
    return &StaffRepository{db: db}
}

func (r *StaffRepository) Create(model *models.Staff) error {
    return r.db.Create(model).Error
}

func (r *StaffRepository) FindByID(id uint) (*models.Staff, error) {
    var model models.Staff
    if err := r.db.First(&model, id).Error; err != nil {
        return nil, err
    }
    return &model, nil
}

func (r *StaffRepository) Update(model *models.Staff) error {
    return r.db.Save(model).Error
}

func (r *StaffRepository) Delete(id uint) error {
    return r.db.Delete(&models.Staff{}, id).Error
}