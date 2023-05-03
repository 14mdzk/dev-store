package schema

type RegistrationReq struct {
	Email    string `json:"email" validate:"required,email"`
	Username string `json:"username" validate:"required,min=6,max=10,alphanum"`
	Password string `json:"password" validate:"required,min=8,alphanum"`
}
