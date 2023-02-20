package merchant

import (
	"Test_MNC_2/domain/merchant/repository"
)

type MerchantServiceInterface interface {
	UpdateBalance(id string, amount int64) error
}

type merchantService struct {
	Repository repository.MerchantRepositoryInterface
}

func MerchantService(repository repository.MerchantRepositoryInterface) MerchantServiceInterface {
	return &merchantService{
		Repository: repository,
	}
}

func (ms *merchantService) UpdateBalance(id string, amount int64) error {
	err := ms.Repository.UpdateBalance(id, amount)
	if err != nil {
		return err
	}
	return nil
}
