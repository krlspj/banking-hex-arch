package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/krlspj/banking-hex-arch/07_security/internal/auth/dto"
	"github.com/krlspj/banking-hex-arch/07_security/internal/auth/service"
)

type AuthHandler struct {
	service service.AuthService
}

func (h AuthHandler) Login(w http.ResponseWriter, r *http.Request) {
	var loginRequest dto.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginRequest); err != nil {
		log.Println("Error while decoding login request" + err.Error())
		w.WriteHeader(http.StatusBadGateway)
	} else {
		token, err := h.service.Login(loginRequest)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprint(w, *token)
		}
	}
}
