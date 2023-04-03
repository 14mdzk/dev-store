package service

import "github.com/14mdzk/dev-store/internal/app/schema"

type ICategoryService interface {
	BrowseAll() ([]schema.GetCategoryResp, error)
	GetById(id int) (schema.GetCategoryResp, error)
	Create(req schema.CreateCategoryReq) error
	Update(id int, req schema.CreateCategoryReq) error
	Delete(id int) error
}

type IUserService interface {
	BrowseAll() ([]schema.GetUserResp, error)
	GetById(id int) (schema.GetUserResp, error)
	Create(req schema.CreateUserReq) error
	Update(id int, req schema.CreateUserReq) error
	Delete(id int) error
}

type IUserAddressService interface {
	BrowseAll(userId int) ([]schema.GetUserAddressResp, error)
	GetById(id int, userId int) (schema.GetUserAddressResp, error)
	Create(userId int, req schema.CreateUserAddressReq) error
	Update(id int, userId int, req schema.CreateUserAddressReq) error
	Delete(id int, userId int) error
}

type IProductService interface {
	BrowseAll() ([]schema.GetProductResp, error)
	GetById(id int) (schema.GetProductResp, error)
	Create(req schema.CreateProductReq) error
	Update(id int, req schema.CreateProductReq) error
	Delete(id int) error
}
