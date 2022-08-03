package mysql

import (
	"database/sql"
	"fmt"

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
