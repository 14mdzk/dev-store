package schema

type GetShoppingCartResp struct {
	ID     int     `json:"id"`
	UserId int     `json:"user_id"`
	Total  float32 `json:"total"`
	Items []GetShoppingCartItemResp `json:"items"`
}

type GetShoppingCartItemResp struct {
	ID        int `json:"id"`
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}

type CreateShoppingCartReq struct {
	UserId int                         `json:"user_id"`
	Total  float32                     `json:"total"`
	Items  []CreateShoppingCartItemReq `json:"items"`
}

type CreateShoppingCartItemReq struct {
	ProductId int `json:"product_id"`
	Quantity  int `json:"quantity"`
}
