package dto

import (
	"net/http"
	"testing"
)

// Test_transactionValidate should return error when transaction type is not deposito ro withdrawl
func TestTransactionValidateType(t *testing.T) {
	// Arrange
	request := TransactionRequest{
		TransactionType: "invalid transaction type",
	}
	// Act
	appError := request.Validate()

	// Assert
	if appError.Message != "Transaction type can only be deposit or withdrawal" {
		t.Error("Invalid message while testing transaction")
	}

	if appError.Code != http.StatusUnprocessableEntity {
		t.Error("Invalid code while testing transaction")
	}
}

func TestTransactionValidateAmount(t *testing.T) {
	request := TransactionRequest{
		Amount:          -3,
		TransactionType: "deposit",
	}

	appError := request.Validate()

	if appError.Message != "Amount cannot be less than zero" {
		t.Error("Invalid error message")
	}
}
