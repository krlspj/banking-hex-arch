package mysql

import (
	"database/sql"
	"errors"
	"log"

	"github.com/krlspj/banking-hex-arch/07_security/internal/auth/domain"
)

type AuthRepositoryDB struct {
	conn *sql.DB
}

func NewAuthRepositoryDB(dbClient *sql.DB) AuthRepositoryDB {
	return AuthRepositoryDB{
		conn: dbClient,
	}
}

func (d AuthRepositoryDB) FindBy(username, password string) (*domain.Login, error) {
	var login domain.Login
	sqlVerify := `SELECT u.username, u.customer_id, role, group_concat(a.account_id) as account_numbers 
				FROM users u
				LEFT JOIN accounts a On a.customer_id = u.customer_id
				WHERE username = ? and password = ?
				GROUP BY u.username
				`
	//GROUP BY a.customer_id` -> producing an error...
	row := d.conn.QueryRow(sqlVerify, username, password)
	err := row.Scan(&login.Username, &login.CustomerId, &login.Role, &login.Accounts)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid credentials")
		} else {

			log.Println("Error while verifying login request from database: " + err.Error())
			return nil, errors.New("Unexpected database error")
		}
	}
	return &login, nil
}
