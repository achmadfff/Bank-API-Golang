package repository

import (
	"errors"
	"gorm.io/gorm"

	"Test_MNC_2/domain/user/model"
)

type RoleRepositoryInterface interface {
	GetRole(offset, limit int) ([]model.Role, int64, error)
}

type roleRepository struct {
	DB *gorm.DB
}

func RoleRepository(db *gorm.DB) RoleRepositoryInterface {
	return &roleRepository{
		DB: db,
	}
}

func (r *roleRepository) GetRole(offset, limit int) ([]model.Role, int64, error) {
	var role []model.Role
	var totalRow int64

	if err := r.DB.Table("roles").Count(&totalRow).Offset(offset).Limit(limit).Scan(&role).Error; err != nil {
		return role, totalRow, errors.New("failed to get role : " + err.Error())
	}

	return role, totalRow, nil
}
