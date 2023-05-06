package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/14mdzk/dev-store/internal/app/model"
	"github.com/14mdzk/dev-store/internal/app/schema"
	"github.com/14mdzk/dev-store/internal/pkg/reason"
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	GetById(id int) (model.User, error)
	GetByEmail(email string) (model.User, error)
}
type AuthRepository interface {
	Create(auth model.Auth) error
	DeleteAllByUserID(userID int) error
	Find(userID int, refreshToken string) (model.Auth, error)
}

type TokenGenerator interface {
	GenerateAccessToken(userID int) (string, time.Time, error)
	GenerateRefreshToken(userID int) (string, time.Time, error)
}

type SessionService struct {
	userRepo   UserRepository
	authRepo   AuthRepository
	tokenMaker TokenGenerator
}

func NewSessionService(
	userRepo UserRepository,
	authRepo AuthRepository,
	tokenMaker TokenGenerator,
) *SessionService {
	return &SessionService{
		userRepo:   userRepo,
		authRepo:   authRepo,
		tokenMaker: tokenMaker,
	}
}

func (ss *SessionService) Login(req *schema.LoginReq) (schema.LoginResp, error) {
	var resp schema.LoginResp

	existingUser, _ := ss.userRepo.GetByEmail(req.Email)
	if existingUser.ID <= 0 {
		return resp, errors.New(reason.LoginFailed)
	}

	isVerified := ss.verifyPassword(existingUser.Password, req.Password)
	if !isVerified {
		return resp, errors.New(reason.LoginFailed)
	}

	accessToken, _, err := ss.tokenMaker.GenerateAccessToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("access token creation : %w", err))
		return resp, errors.New(reason.LoginFailed)
	}

	refreshToken, expiredAt, err := ss.tokenMaker.GenerateRefreshToken(existingUser.ID)
	if err != nil {
		log.Error(fmt.Errorf("refresh token creation: %w", err))
		return resp, errors.New(reason.LoginFailed)
	}

	resp.AccessToken = accessToken
	resp.RefreshToken = refreshToken

	authPayload := model.Auth{
		UserID:    existingUser.ID,
		Token:     refreshToken,
		AuthType:  "refresh_token",
		ExpiredAt: expiredAt,
	}

	err = ss.authRepo.Create(authPayload)
	if err != nil {
		log.Error(fmt.Errorf("refresh token saving: %w", err))
		return resp, errors.New(reason.LoginFailed)
	}

	return resp, nil
}
func (ss *SessionService) Logout(userID int) error {
	err := ss.authRepo.DeleteAllByUserID(userID)
	if err != nil {
		log.Error(fmt.Errorf("remove refresh token: %w", err))
		return err
	}

	return nil
}
func (ss *SessionService) Refresh(req *schema.RefreshTokenReq) (schema.RefreshTokenResp, error) {
	var resp schema.RefreshTokenResp
	log.Error(fmt.Print(req))
	existingUser, _ := ss.userRepo.GetById(req.UserID)
	if existingUser.ID <= 0 {
		return resp, errors.New(reason.RefreshTokenFailed)
	}

	auth, err := ss.authRepo.Find(req.UserID, req.RefreshToken)
	if err != nil || auth.UserID <= 0 {
		log.Error(fmt.Errorf("error SessionService - Refresh: %w", err))
		return resp, errors.New(reason.RefreshTokenFailed)
	}

	accessToken, _, _ := ss.tokenMaker.GenerateAccessToken(existingUser.ID)
	resp.AccessToken = accessToken
	return resp, nil
}

func (ss *SessionService) verifyPassword(hashedPassword string, plainPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	return err == nil
}
