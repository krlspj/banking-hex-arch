package domain

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type CustomerRepositoryDB struct {
	conn *sql.DB
}

func NewCustomerRepositoryDB() CustomerRepositoryDB {
	//client, err := sql.Open("mysql", "user:password@/dbname")
	client, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/banking")
	if err != nil {
		panic(err)
	}

	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDB{
		conn: client,
	}
}

func (cdb CustomerRepositoryDB) FindAll() ([]Customer, error) {

	findAllSql := "select customer_id, name, city, zipcode, date_of_birth, status from customers"

	rows, err := cdb.conn.Query(findAllSql)

	if err != nil {
		log.Println("Error while reading customers table:", err.Error())
		return []Customer{}, err
	}

	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.Zipcode, &c.DateOfBirth, &c.Status)
		if err != nil {
			log.Println("Error while scanning customers table:", err.Error())
			return []Customer{}, err
		}
		customers = append(customers, c)
	}

	return customers, nil

}
