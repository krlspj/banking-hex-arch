package mysql

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/domain"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/errs"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/logger"
)

type AccountRepositoryDB struct {
	conn *sql.DB
}

func NewAccountRepositoryDB(dbClient *sql.DB) AccountRepositoryDB {
	return AccountRepositoryDB{
		conn: dbClient,
	}
}

func (d AccountRepositoryDB) Save(a domain.Account) (*domain.Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts (customer_id, opening_date, account_type, amount, status) values (?, ?, ?, ?, ?)"
	result, err := d.conn.Exec(sqlInsert, a.CustomerId, a.OpeningDate, a.AccountType, a.Amount, a.Status)
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpedtedError(err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		logger.Error(err.Error())
		return nil, errs.NewUnexpedtedError(err.Error())
	}
	//a.AccountId = strconv.FormatInt(id, 10)
	a.AccountId = fmt.Sprint(id)
	return &a, nil
}

func (d AccountRepositoryDB) SaveTransaction(t domain.Transaction) (*domain.Transaction, *errs.AppError) {
	// starting the database transaction block

	tx, err := d.conn.Begin()
	if err != nil {
		logger.Error("Error while starting a new transaction for back account transaction: " + err.Error())
		return nil, errs.NewUnexpedtedError("Unexpected database error")
	}

	// inserting bank account transaction
	result, _ := tx.Exec(`INSERT INTO transactions (account_id, amount, transaction_type, transaction_date) 
	values (?, ?, ?, ?)`, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)

	// updating account balance
	if t.IsWithrawal() {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount - ? where account_id = ?`,
			t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`UPDATE accounts SET amount = amount + ? where account_id = ?`,
			t.Amount, t.AccountId)
	}

	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving transaction: " + err.Error())
		return nil, errs.NewUnexpedtedError("Unexpected database error")
	}
	// commit the transaction when all is good
	err = tx.Commit()
	if err != nil {
		logger.Error("Error while commintng transaction for bank account: " + err.Error())
		return nil, errs.NewUnexpedtedError("Unexpected database error")
	}

	// getting the last transaction ID from the transaction table
	transactionId, err := result.LastInsertId()
	if err != nil {
		logger.Error("Error while getting the last transaction id: " + err.Error())
		return nil, errs.NewUnexpedtedError("Unexpected database error")
	}

	// getting the lastest account information from the accounts table
	account, appErr := d.FindById(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)

	// updating the transaction struct with the latest balance
	t.Amount = account.Amount
	return &t, nil
}

func (d AccountRepositoryDB) FindById(accountId string) (*domain.Account, *errs.AppError) {
	accountFindById := "SELECT account_id, customer_id, opening_date, account_type, amount, status " +
		"FROM accounts WHERE account_id = ?"
	row := d.conn.QueryRow(accountFindById, accountId)

	var a domain.Account
	err := row.Scan(&a.AccountId, &a.CustomerId, &a.OpeningDate, &a.AccountType, &a.Amount, &a.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpedtedError("Unexpected database error")
		}
	}
	return &a, nil
}
