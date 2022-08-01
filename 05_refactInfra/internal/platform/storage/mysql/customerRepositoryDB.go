package mysql

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/krlspj/banking-hex-arch/05_refactInfra/internal/domain"
)

type CustomerRepositoryDB struct {
	conn *sql.DB
}

func NewCustomerRepositoryDB() *CustomerRepositoryDB {
	//client, err := sql.Open("mysql", "user:password@/dbname")
	client, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return &CustomerRepositoryDB{
		conn: client,
	}
}

func (cdb *CustomerRepositoryDB) FindAll() ([]domain.Customer, error) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := cdb.conn.Query(findAllSql)

	if err != nil {
		log.Println("Error while reading customers table:", err.Error())
		return []domain.Customer{}, err
	}

	customers := make([]domain.Customer, 0)
	for rows.Next() {
		var c domain.Customer
		err := rows.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			if err == sql.ErrNoRows {
				return nil, errors.New("Customers not found")
			} else {
				log.Println("Error while scanning customers table:", err.Error())
				return []domain.Customer{}, errors.New("Unexpected database error")
			}
		}
		customers = append(customers, c)
	}
	return customers, nil
}

func (cbd *CustomerRepositoryDB) ById(id string) (*domain.Customer, error) {
	customerSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers where customer_id = ? "
	row := cbd.conn.QueryRow(customerSql, id)
	var c domain.Customer
	err := row.Scan(&c.ID, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("Customer not found")
		} else {
			log.Println("Error while scanning customer " + err.Error())
			return nil, errors.New("Unexpected database error")
		}
	}
	return &c, nil
}
