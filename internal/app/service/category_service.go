package service

import (
	"errors"

	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/14mdzk/dev-store/internal/app/repository"
	"github.com/14mdzk/dev-store/internal/app/schema"
	"github.com/14mdzk/dev-store/internal/pkg/reason"
)

type CategoryService struct {
	repo repository.ICategoryRepository
}

func NewCategoryService(repo repository.ICategoryRepository) *CategoryService {
	return &CategoryService{
		repo: repo,
	}
}

func (cs *CategoryService) BrowseAll() ([]schema.GetCategoryResp, error) {
	var response []schema.GetCategoryResp

	categories, err := cs.repo.Browse()
	if err != nil {
		return nil, errors.New(reason.CategoryNotFound)
	}

	for _, value := range categories {
		var data schema.GetCategoryResp
		data.ID = value.ID
		data.Name = value.Name
		data.Description = value.Description

		response = append(response, data)
	}

	return response, nil
}

func (cs *CategoryService) GetById(id int) (schema.GetCategoryResp, error) {
	var response schema.GetCategoryResp

	category, err := cs.repo.GetById(id)

	if err != nil {
		return response, err
	}

	response.ID = category.ID
	response.Name = category.Name
	response.Description = category.Description

	return response, nil
}

func (cs *CategoryService) Create(req schema.CreateCategoryReq) error {
	var category model.Category
	category.Name = req.Name
	category.Description = req.Description

	err := cs.repo.Create(category)
	if err != nil {
		return errors.New(reason.CategoryCreateFailed)
	}

	return nil
}

func (cs *CategoryService) Update(id int, req schema.CreateCategoryReq) error {
	var category model.Category
	category.ID = id
	category.Name = req.Name
	category.Description = req.Description

	err := cs.repo.Update(category)
	if err != nil {
		return errors.New(reason.CategoryUpdateFailed)
	}

	return nil
}

func (cs *CategoryService) Delete(id int) error {
	err := cs.repo.Delete(id)
	if err != nil {
		return errors.New(reason.CategoryDeleteFailed)
	}

	return nil
}
