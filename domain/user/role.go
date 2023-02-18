package user

import (
	"net/http"

	"Test_MNC_2/domain/user/model"
	"Test_MNC_2/domain/user/repository"
)

type RoleServiceInterface interface {
	GetRole(offset, limit int) ([]model.Role, int64, int, error)
}

type roleService struct {
	Repository repository.RoleRepositoryInterface
}

func RoleService(repository repository.RoleRepositoryInterface) RoleServiceInterface {
	return &roleService{
		Repository: repository,
	}
}

func (s *roleService) GetRole(offset, limit int) ([]model.Role, int64, int, error) {
	var roles []model.Role
	roles, totalRow, err := s.Repository.GetRole(offset, limit)
	if err != nil {
		return roles, totalRow, http.StatusInternalServerError, err
	}

	return roles, totalRow, http.StatusOK, nil
}
