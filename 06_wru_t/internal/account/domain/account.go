package domain

import (
	"github.com/krlspj/banking-hex-arch/06_wru_t/internal/account/dto"
	"github.com/krlspj/banking-hex-arch/06_wru_t/internal/errs"
)

type Account struct {
	AccountId   string
	CustomerId  string
	OpeningDate string
	AccountType string
	Amount      float64
	Status      string
}

func (a Account) ToNewAccountResponseDto() dto.NewAccountResponse {
	return dto.NewAccountResponse{
		AccountId: a.AccountId,
	}
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
}
