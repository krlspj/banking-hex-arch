package app

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/krlspj/banking-hex-arch/06_wru_t/internal/account/dto"
	"github.com/krlspj/banking-hex-arch/06_wru_t/internal/account/service"
)

type AccountHandlers struct {
	service service.AccountService
}

func (ah AccountHandlers) CreateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	customerId := vars["customer_id"]
	var request dto.NewAccountRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		writeResponse(w, http.StatusBadRequest, err.Error())
	} else {
		request.CustomerId = customerId
		account, appError := ah.service.NewAccount(request)
		if appError != nil {
			writeResponse(w, appError.Code, appError.Message)
		} else {
			writeResponse(w, http.StatusCreated, account)
		}
	}
}
