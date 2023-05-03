package service

import (
	"errors"

	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/14mdzk/dev-store/internal/app/repository"
	"github.com/14mdzk/dev-store/internal/app/schema"
	"github.com/14mdzk/dev-store/internal/pkg/reason"
	"golang.org/x/crypto/bcrypt"
)

type RegistrationService struct {
	repo repository.IUserRepository
}

func NewRegistrationService(repo repository.IUserRepository) *RegistrationService {
	return &RegistrationService{
		repo: repo,
	}
}

func (rs *RegistrationService) Register(req schema.RegistrationReq) error {
	// check existing user
	existingUser, _ := rs.repo.GetByEmail(req.Email)

	if existingUser.ID > 0 {
		return errors.New(reason.UserAlreadyExist)
	}

	password, _ := rs.hashPassword(req.Password)

	var user model.User
	user.Username = req.Username
	user.Email = req.Email
	user.Password = password
	err := rs.repo.Create(user)

	if err != nil {
		return errors.New(reason.RegistrationFailed)
	}

	return nil
}

func (rs *RegistrationService) hashPassword(password string) (string, error) {
	bytePassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytePassword), nil
}
