package schema

type GetUserResp struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type CreateUserReq struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type GetUserAddressResp struct {
	ID          int    `json:"id"`
	UserId      int    `json:"user_id"`
	AddressLine string `json:"address_line"`
	Country     string `json:"country"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	Phone       string `json:"phone"`
	Note        string `json:"note"`
	IsMain      bool   `json:"is_main"`
}

type CreateUserAddressReq struct {
	AddressLine string `json:"address_line"`
	Country     string `json:"country"`
	City        string `json:"city"`
	PostalCode  string `json:"postal_code"`
	Phone       string `json:"phone"`
	Note        string `json:"note"`
	IsMain      bool   `json:"is_main"`
}
