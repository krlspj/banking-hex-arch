package dto

import (
	"strings"

	"github.com/krlspj/banking-hex-arch/06_wru_t/internal/errs"
)

type NewAccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (r NewAccountRequest) Validate() *errs.AppError {
	if r.Amount < 5000 {
		return errs.NewAccountValidationError("To Open a new account you need to deposit at least 5000")
	}
	if strings.ToLower(r.AccountType) != "saving" && strings.ToLower(r.AccountType) != "checking" {
		return errs.NewAccountValidationError("Account type must be checking or saving")
	}
	return nil
}
