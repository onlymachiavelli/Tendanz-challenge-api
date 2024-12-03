package requests

type ClientRegister struct {
	FirstName string `json:"first_name" form:"first_name" query:"first_name" validate:"required"`
	LastName  string `json:"last_name" form:"last_name" query:"last_name" validate:"required"`
	Email     string `json:"email" form:"email" query:"email" validate:"required,email"`
	Password  string `json:"password" form:"password" query:"password" validate:"required"`
	Phone     string `json:"phone" form:"phone" query:"phone" validate:"required"`
}

type LoginClient struct {
	Email    string `json:"email" form:"email" query:"email" validate:"required,email"`
	Password string `json:"password" form:"password" query:"password" validate:"required"`
}

type VerifyCode struct {
	Code string `json:"code" form:"code" query:"code" validate:"required"`
}