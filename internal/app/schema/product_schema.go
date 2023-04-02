package schema

type GetProductResp struct {
	ID          int     `json:"id"`
	CategoryId  int     `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Currency    string  `json:"currency"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
	IsActive    bool    `json:"is_active"`
}

type CreateProductReq struct {
	CategoryId  int     `json:"category_id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Currency    string  `json:"currency"`
	Price       float32 `json:"price"`
	Stock       int     `json:"stock"`
	IsActive    bool    `json:"is_active"`
}
