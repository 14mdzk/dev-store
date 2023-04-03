package service

import (
	"errors"

	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/14mdzk/dev-store/internal/app/repository"
	"github.com/14mdzk/dev-store/internal/app/schema"
)

type UserAddressService struct {
	repo     repository.IUserAddressRepository
	userRepo repository.IUserRepository
}

func NewUserAddressService(repo repository.IUserAddressRepository, userRepo repository.IUserRepository) *UserAddressService {
	return &UserAddressService{
		repo:     repo,
		userRepo: userRepo,
	}
}

func (cs *UserAddressService) BrowseAll(userId int) ([]schema.GetUserAddressResp, error) {
	var response []schema.GetUserAddressResp

	users, err := cs.repo.Browse(userId)
	if err != nil {
		return nil, errors.New("cannot get user addresses")
	}

	for _, userAddress := range users {
		var data schema.GetUserAddressResp
		data.ID = userAddress.ID
		data.UserId = userAddress.UserId
		data.AddressLine = userAddress.AddressLine
		data.Country = userAddress.Country
		data.City = userAddress.City
		data.PostalCode = userAddress.PostalCode
		data.Phone = userAddress.Phone
		data.Note = userAddress.Note
		data.IsMain = userAddress.IsMain

		response = append(response, data)
	}

	return response, nil
}

func (cs *UserAddressService) GetById(id int, userId int) (schema.GetUserAddressResp, error) {
	var response schema.GetUserAddressResp

	userAddress, err := cs.repo.GetById(id, userId)

	if err != nil {
		return response, err
	}

	response.ID = userAddress.ID
	response.UserId = userAddress.UserId
	response.AddressLine = userAddress.AddressLine
	response.Country = userAddress.Country
	response.City = userAddress.City
	response.PostalCode = userAddress.PostalCode
	response.Phone = userAddress.Phone
	response.Note = userAddress.Note
	response.IsMain = userAddress.IsMain

	return response, nil
}

func (cs *UserAddressService) Create(userId int, req schema.CreateUserAddressReq) error {
	user, err := cs.userRepo.GetById(userId)
	if err != nil {
		return errors.New("cannot create user address, user not found")
	}

	var userAddress model.UserAddress

	userAddress.UserId = user.ID
	userAddress.AddressLine = req.AddressLine
	userAddress.Country = req.Country
	userAddress.City = req.City
	userAddress.PostalCode = req.PostalCode
	userAddress.Phone = req.Phone
	userAddress.Note = req.Note
	userAddress.IsMain = req.IsMain

	err = cs.repo.Create(userAddress)
	if err != nil {
		return errors.New("cannot create user address")
	}

	return nil
}

func (cs *UserAddressService) Update(id int, userId int, req schema.CreateUserAddressReq) error {
	var userAddress model.UserAddress

	userAddress.ID = id
	userAddress.UserId = userId
	userAddress.AddressLine = req.AddressLine
	userAddress.Country = req.Country
	userAddress.City = req.City
	userAddress.PostalCode = req.PostalCode
	userAddress.Phone = req.Phone
	userAddress.Note = req.Note
	userAddress.IsMain = req.IsMain

	err := cs.repo.Update(userAddress)
	if err != nil {
		return errors.New("cannot update user address")
	}

	return nil
}

func (cs *UserAddressService) Delete(id int, userId int) error {
	err := cs.repo.Delete(id, userId)
	if err != nil {
		return errors.New("cannot delete user address")
	}

	return nil
}
