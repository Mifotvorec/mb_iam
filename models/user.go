package models

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

const solt = "solty" //save solt like a const, better to use ENV

// User structure
type User struct {
	Login      string `json:"login"`
	Name       string `json:"name"`
	MiddleName string `json:"middlename"`
	Surname    string `json:"surname"`
	//	Status     byte      `json:"status"`
	Email      string    `json:"email"`
	Phone      string    `json:"phone"`
	Passwd     string    `json:"passwd"`
	DateCreate time.Time `json:"datecreate"`
	DateUpdate time.Time `json:"dateupdate"`
	DateDelete time.Time `json:"datedelete"`
	//ID         uuid.UUID `json:"id"`
}

// Function save one User to database
func userCreateNew(c *User) error {

	hash, _ := passwdMakeHashe(c.Passwd) // ignore error for the sake of simplicity
	stmt, err := database.Prepare(`
		INSERT INTO users_main (login, name, middleaname, surname, email, phone, passwd, datecreate)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.Login, c.Name, c.MiddleName, c.Surname, c.Email, c.Phone, hash, time.Now())
	if err != nil {
		return err
	}

	return nil
}

// Function update one User in database
func userUpdateByLogin(c *User) error {

	stmt, err := database.Prepare(`
		UPDATE users_main SET name = $2, middleaname = $3, surname = $4, email = $6, phone = $7, dateupdate = $8
		WHERE login = $1 
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.Login, c.Name, c.MiddleName, c.Surname, c.Email, c.Phone, time.Now())
	if err != nil {
		return err
	}
	return nil

}

// Function delete one User in database (mark as delete)
func userDeleteByLogin(c *User) error {
	stmt, err := database.Prepare(`
	UPDATE users_main SET datedelete = $2
	WHERE login = $1
`)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(c.Login, time.Now())
	if err != nil {
		return err
	}
	return nil
}

// Function read one User in database (full raw)
func userGetByLogin(login string, passwd string) (*User, error) {

	hash, _ := passwdMakeHashe(passwd) // ignore error for the sake of simplicity

	stmt, err := database.Prepare(`
		SELECT login, name, middlename, surname, status, email, phone, datecreate, dateupdate, datedelete
		FROM users_main
		WHERE login = $1 and passwd = $2
	`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	row := stmt.QueryRow(login, hash)
	user := &User{}
	err = row.Scan(
		&user.Login,
		&user.Name,
		&user.MiddleName,
		&user.Surname,
		&user.Email,
		&user.Phone,
		&user.DateCreate,
		&user.DateUpdate,
		&user.DateDelete,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user with login %s not found", login)
		}
		return nil, err
	}

	return user, nil
}

// Function for hashing password with adding solt.
// solt using like a Const
func passwdMakeHashe(passwd string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(passwd+solt), 15)
	return string(bytes), err
}

func passwdCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password+solt))
	return err == nil
}
