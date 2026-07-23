package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"spotsync/internal/dto"
	"spotsync/internal/models"
	"spotsync/internal/repository"
	"spotsync/internal/auth"
)

type AuthService struct {
	repo repository.UserRepository
}

func NewAuthService(repo repository.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error) {

	_, err := s.repo.GetByEmail(req.Email)

	if err == nil {
		return nil, errors.New("email already exists")
	}

	hash, err := bcrypt.GenerateFromPassword(
		[]byte(req.Password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		return nil, err
	}

	role := req.Role
	if role == "" {
		role = "driver"
	}

	user := &models.User{
		Name: req.Name,
		Email: req.Email,
		Password: string(hash),
		Role: role,
	}

	if err := s.repo.Create(user); err != nil {
		return nil, err
	}

	return &dto.RegisterResponse{
		ID: user.ID,
		Name: user.Name,
		Email: user.Email,
		Role: user.Role,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
}

func (s *AuthService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {

	user, err := s.repo.GetByEmail(req.Email)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(req.Password),
	)

	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	token, err := auth.GenerateJWT(user.ID, user.Role)

	if err != nil {
		return nil, err
	}

	return &dto.LoginResponse{
		Token: token,
		User: dto.UserResponse{
			ID: user.ID,
			Name: user.Name,
			Email: user.Email,
			Role: user.Role,
		},
	}, nil
}