package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/krlspj/banking-hex-arch/06_wru_t/internal/customer/domain"
	"github.com/krlspj/banking-hex-arch/06_wru_t/internal/errs"
	"github.com/krlspj/banking-hex-arch/06_wru_t/internal/logger"
)

type CustomerRepositoryDB struct {
	conn *sql.DB
}

func NewCustomerRepositoryDB(dbClient *sql.DB) *CustomerRepositoryDB {
	return &CustomerRepositoryDB{
		conn: dbClient,
	}
}

func (cdb *CustomerRepositoryDB) FindAll(status string) ([]domain.Customer, *errs.AppError) {
	var rows *sql.Rows
	var err error
	if status == "" {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"
		rows, err = cdb.conn.Query(findAllSql)
	} else {
		findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where status = ?"
		rows, err = cdb.conn.Query(findAllSql, status)
	}

	if err != nil {
		log.Println("Error while reading customers table:", err.Error())
		return nil, errs.NewUnexpedtedError("Unexpected database error")
	}

	customers := make([]domain.Customer, 0)
	for rows.Next() {
		var c domain.Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			if err == sql.ErrNoRows {
				logger.Error("Error while scanning customers " + err.Error())
				return nil, errs.NewNotFoundError("Customers not found")
			} else {
				logger.Error("Error while scanning customers table:" + err.Error())
				return nil, errs.NewUnexpedtedError("Unexpected database error")
			}
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (cbd *CustomerRepositoryDB) ById(id string) (*domain.Customer, *errs.AppError) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ? "
	row := cbd.conn.QueryRow(customerSql, id)
	var c domain.Customer
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning customer " + err.Error())
			return nil, errs.NewUnexpedtedError("Unexpected database error")
		}
	}
	return &c, nil
}
