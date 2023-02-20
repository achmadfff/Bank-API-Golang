package repository

import (
	"gorm.io/gorm"

	"Test_MNC_2/domain/user/model"
)

type AuthRepositoryInterface interface {
	FirstByEmail(email *string) (model.User, error)
	FirstByID(string) (model.User, error)
	Create(reqBody *model.ReqBodyRegister) error
	UpdateToken(model.User, string) error
	DeleteToken(model.User) error
}

type authRepository struct {
	DB *gorm.DB
}

func AuthRepository(DB *gorm.DB) AuthRepositoryInterface {
	return &authRepository{
		DB: DB,
	}
}

func (r *authRepository) FirstByEmail(email *string) (user model.User, err error) {
	if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *authRepository) FirstByID(id string) (model.User, error) {
	var user model.User
	if err := r.DB.Preload("Role").Where("id = ?", id).First(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *authRepository) Create(reqBody *model.ReqBodyRegister) error {
	var user model.User
	user.Name = reqBody.Name
	user.Email = reqBody.Email
	user.EncryptedPassword = reqBody.Password
	if reqBody.RoleID == "" {
		user.RoleID = "1ca78238-8e40-4763-a20f-59b5b41791b2"
	}
	return r.DB.Create(&user).Error
}

func (r *authRepository) UpdateToken(user model.User, ss string) error {
	user.Token = ss
	return r.DB.Save(&user).Error
}

func (r *authRepository) DeleteToken(user model.User) error {
	user.Token = ""
	return r.DB.Save(&user).Error
}
