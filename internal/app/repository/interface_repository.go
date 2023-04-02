package repository

import "github.com/14mdzk/dev-store/internal/app/model"

type ICategoryRepository interface {
	Browse() ([]model.Category, error)
	GetById(id int) (model.Category, error)
	Create(model.Category) error
	Update(model.Category) error
	Delete(id int) error
}

type IUserRepository interface {
	Browse() ([]model.User, error)
	GetById(id int) (model.User, error)
	Create(model.User) error
	Update(model.User) error
	Delete(id int) error
}
