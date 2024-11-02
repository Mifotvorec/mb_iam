package models

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/lib/pq"
)

// User structure
type Account struct {
	Account string `json:"account"`
	//	Type       string    `json:"type"`
	Login      string    `json:"login"`
	Balance    int       `json:"balance"`
	DateCreate time.Time `json:"datecreate"`
	DateUpdate time.Time `json:"dateupdate"`
	DateDelete time.Time `json:"datedelete"`
	//ID         uuid.UUID `json:"id"`
}

// Function create one Account to database
func accountCreateNew(a *Account) error {

	stmt, err := database.Prepare(`
		INSERT INTO accounts (account, login, datecreate)
		VALUES ($1, $2, $3)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(a.Account, a.Login, time.Now())
	if err != nil {
		return err
	}

	return nil
}

// Function get Accounts by User
func accountGetListbyUser(login string) ([]Account, error) {
	var accounts []Account
	rows, err := database.Query("SELECT account, login, balance, datecreate, dateupdate  FROM accounts where login =$1 and datedelete is NULL")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var acc Account
		err := rows.Scan(&acc.Account, &acc.Login, &acc.Balance, &acc.DateCreate, &acc.DateUpdate)
		if err != nil {
			log.Fatal(err)
		}
		accounts = append(accounts, acc)
	}
	return accounts, err
}

// Function read one Account from database (full raw)
func accountGetByNum(accountNum string) (*Account, error) {

	stmt, err := database.Prepare(`
		SELECT account, login, balance, datecreate, dateupdate 
		FROM accounts
		WHERE account = $1 and datedelete is NULL
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(accountNum)
	account := &Account{}

	err = row.Scan(
		&account.Account,
		&account.Login,
		account.Balance,
		&account.DateCreate,
		&account.DateUpdate,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("Account %s not found", accountNum)
		}
		return nil, err
	}

	return account, nil
}
