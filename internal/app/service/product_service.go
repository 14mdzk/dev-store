package service

import (
	"errors"

	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/14mdzk/dev-store/internal/app/repository"
	"github.com/14mdzk/dev-store/internal/app/schema"
)

type ProductService struct {
	repo repository.IProductRepository
}

func NewProductService(repo repository.IProductRepository) *ProductService {
	return &ProductService{
		repo: repo,
	}
}

func (cs *ProductService) BrowseAll() ([]schema.GetProductResp, error) {
	var response []schema.GetProductResp

	products, err := cs.repo.Browse()
	if err != nil {
		return nil, errors.New("cannot get products")
	}

	for _, product := range products {
		var data schema.GetProductResp
		data.ID = product.ID
		data.CategoryId = product.CategoryId
		data.Name = product.Name
		data.Description = product.Description
		data.Currency = product.Currency
		data.Price = product.Price
		data.Stock = product.Stock
		data.IsActive = product.IsActive

		response = append(response, data)
	}

	return response, nil
}

func (cs *ProductService) GetById(id int) (schema.GetProductResp, error) {
	var response schema.GetProductResp

	product, err := cs.repo.GetById(id)

	if err != nil {
		return response, err
	}

	response.ID = product.ID
	response.CategoryId = product.CategoryId
	response.Name = product.Name
	response.Description = product.Description
	response.Currency = product.Currency
	response.Price = product.Price
	response.Stock = product.Stock
	response.IsActive = product.IsActive

	return response, nil
}

func (cs *ProductService) Create(req schema.CreateProductReq) error {
	var product model.Product
	product.CategoryId = req.CategoryId
	product.Name = req.Name
	product.Description = req.Description
	product.Currency = req.Currency
	product.Price = req.Price
	product.Stock = req.Stock
	product.IsActive = req.IsActive

	err := cs.repo.Create(product)
	if err != nil {
		return errors.New("cannot create product")
	}

	return nil
}

func (cs *ProductService) Update(id int, req schema.CreateProductReq) error {
	var product model.Product
	product.ID = id
	product.CategoryId = req.CategoryId
	product.Name = req.Name
	product.Description = req.Description
	product.Currency = req.Currency
	product.Price = req.Price
	product.Stock = req.Stock
	product.IsActive = req.IsActive

	err := cs.repo.Update(product)
	if err != nil {
		return errors.New("cannot update product")
	}

	return nil
}

func (cs *ProductService) Delete(id int) error {
	err := cs.repo.Delete(id)
	if err != nil {
		return errors.New("cannot delete Product")
	}

	return nil
}
