package requests

type REGISTERADMINREQUEST struct {
	Email       string `json:"email" form:"email" query:"email" validate:"required,email"`
	Password    string `json:"password" form:"password" query:"password" validate:"required"`
	Identity    string `json:"identity" form:"identity" query:"identity" validate:"required"`
	FirstName   string `json:"first_name" form:"first_name" query:"first_name" validate:"required"`
	LastName    string `json:"last_name" form:"last_name" query:"last_name" validate:"required"`
	PhoneNumber string `json:"phone_number" form:"phone_number" query:"phone_number" validate:"required"`
}

type LOGINADMIN struct {
	Email    string `json:"email" form:"email" query:"email" validate:"required,email"`
	Password string `json:"password" form:"password" query:"password" validate:"required"`
}