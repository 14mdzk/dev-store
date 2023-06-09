package repository

import "github.com/14mdzk/dev-store/internal/app/model"

type IUserRepository interface {
	Browse() ([]model.User, error)
	GetById(id int) (model.User, error)
	GetByEmail(email string) (model.User, error)
	Create(model.User) error
	Update(model.User) error
	Delete(id int) error
}

type IUserAddressRepository interface {
	Browse(userId int) ([]model.UserAddress, error)
	GetById(id int, userId int) (model.UserAddress, error)
	Create(model.UserAddress) error
	Update(model.UserAddress) error
	Delete(id int, userId int) error
}

type IProductRepository interface {
	Browse() ([]model.Product, error)
	GetById(id int) (model.Product, error)
	Create(model.Product) error
	Update(model.Product) error
	Delete(id int) error
}
