package service

import (
	"errors"

	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/14mdzk/dev-store/internal/app/repository"
	"github.com/14mdzk/dev-store/internal/app/schema"
)

type UserService struct {
	repo repository.IUserRepository
}

func NewUserService(repo repository.IUserRepository) *UserService {
	return &UserService{
		repo: repo,
	}
}

func (cs *UserService) BrowseAll() ([]schema.GetUserResp, error) {
	var response []schema.GetUserResp

	users, err := cs.repo.Browse()
	if err != nil {
		return nil, errors.New("cannot get users")
	}

	for _, value := range users {
		var data schema.GetUserResp
		data.ID = value.ID
		data.Username = value.Username
		data.Email = value.Email

		response = append(response, data)
	}

	return response, nil
}

func (cs *UserService) GetById(id int) (schema.GetUserResp, error) {
	var response schema.GetUserResp

	user, err := cs.repo.GetById(id)

	if err != nil {
		return response, err
	}

	response.ID = user.ID
	response.Username = user.Username
	response.Email = user.Email

	return response, nil
}

func (cs *UserService) Create(req schema.CreateUserReq) error {
	var user model.User
	user.Username = req.Username
	user.Email = req.Email
	user.Password = req.Password

	err := cs.repo.Create(user)
	if err != nil {
		return errors.New("cannot create User")
	}

	return nil
}

func (cs *UserService) Update(id int, req schema.CreateUserReq) error {
	user, err := cs.repo.GetById(id)
	if err != nil {
		return err
	}

	if req.Username != "" {
		user.Username = req.Username
	}

	if req.Password != "" {
		user.Password = req.Password
	}

	if req.Email != "" {
		user.Email = req.Email
	}

	err = cs.repo.Update(user)
	if err != nil {
		return errors.New("cannot update User")
	}

	return nil
}

func (cs *UserService) Delete(id int) error {
	err := cs.repo.Delete(id)
	if err != nil {
		return errors.New("cannot delete User")
	}

	return nil
}
