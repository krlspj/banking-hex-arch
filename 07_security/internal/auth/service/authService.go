package service

import (
	"github.com/krlspj/banking-hex-arch/07_security/internal/auth/domain"
	"github.com/krlspj/banking-hex-arch/07_security/internal/auth/dto"
)

type AuthService interface {
	Login(dto.LoginRequest) (*string, error)
}

type DefaultAuthService struct {
	repo domain.AuthRepository
}

func (s DefaultAuthService) Login(req dto.LoginRequest) (*string, error) {
	login, err := s.repo.FindBy(req.Username, req.Password)
	if err != nil {
		return nil, err
	}
	token, err := login.GenerateToken()
	if err != nil {
		return nil, err
	}
	return token, nil
}

func NewAuthService(repo domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{
		repo: repo,
	}
}
