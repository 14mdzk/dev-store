package service

import (
	"github.com/14mdzk/dev-store/internal/app/repository"
	"github.com/14mdzk/dev-store/internal/app/schema"
)

type ShoppingCartService struct {
	repo repository.IShoppingCartRepository
}

func NewShoppingCartService(repo repository.IShoppingCartRepository) *ShoppingCartService {
	return &ShoppingCartService{
		repo: repo,
	}
}

func (scs *ShoppingCartService) Browse(userId int) (schema.GetShoppingCartResp, error) {
	var (
		response     schema.GetShoppingCartResp
		responseItem []schema.GetShoppingCartItemResp
	)

	cart, err := scs.repo.Browse(userId)
	if err != nil {
		return response, err
	}

	for _, value := range cart.Items {
		var item schema.GetShoppingCartItemResp
		item.ID = value.ID
		item.ProductId = value.ProductId
		item.Quantity = value.ShoppingCartId

		responseItem = append(responseItem, item)
	}

	response.ID = cart.ID
	response.UserId = cart.UserId
	response.Items = responseItem
	response.Total = cart.Total

	return response, nil
}
