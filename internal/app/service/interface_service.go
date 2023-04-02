package service

import "github.com/14mdzk/dev-store/internal/app/schema"

type ICategoryService interface {
	BrowseAll() ([]schema.GetCategoryResp, error)
	GetById(id int) (schema.GetCategoryResp, error)
	Create(req schema.CreateCategoryReq) error
	Update(id int, req schema.CreateCategoryReq) error
	Delete(id int) error
}
