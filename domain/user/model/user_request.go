package model

type ReqBodyRegister struct {
	RoleID   string
	Name     string `binding:"required"`
	Email    string `binding:"required,email"`
	Password string `binding:"required,gte=6"`
}

type ReqBodyLogin struct {
	Email    string `binding:"required,email"`
	Password string `binding:"required,gte=6"`
}
