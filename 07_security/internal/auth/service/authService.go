package service

import (
	"encoding/json"
	"log"
	"net/http"
	"net/url"

	"github.com/krlspj/banking-hex-arch/07_security/internal/auth/domain"
	"github.com/krlspj/banking-hex-arch/07_security/internal/auth/dto"
	"github.com/krlspj/banking-hex-arch/07_security/internal/logger"
)

type AuthService interface {
	Login(dto.LoginRequest) (*string, error)
	IsAuthorized(token, routeName string, vars map[string]string) bool
}

type DefaultAuthService struct {
	repo domain.AuthRepository
}

func NewAuthService(repo domain.AuthRepository) DefaultAuthService {
	return DefaultAuthService{
		repo: repo,
	}
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

func (s DefaultAuthService) IsAuthorized(token, routeName string, vars map[string]string) bool {
	u := s.buildVerifyURL(token, routeName, vars)
	if response, err := http.Get(u); err != nil {
		log.Println("Error on IsAuthorized:", err.Error())
		return false
	} else {
		m := make(map[string]bool)
		if err = json.NewDecoder(response.Body).Decode(&m); err != nil {
			logger.Error("Errow while decoding resopnse from auth server:" + err.Error())
			return false
		}

		return m["isAuthorized"]
	}
}

func (s DefaultAuthService) buildVerifyURL(token, routeName string, vars map[string]string) string {
	u := url.URL{Host: "localhost:8000", Path: "/auth/verify", Scheme: "http"}
	q := u.Query()
	q.Add("token", token)
	q.Add("routeName", routeName)
	for k, v := range vars {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()
	return u.String()
}
