package models

import (
	"errors"

	"github.com/okwu-john/webapi/db"
	"github.com/okwu-john/webapi/utils"
)

type User struct {
	ID       int64
	Email    string `binding: required`
	Password string `binding: required`
}

func (u User) Createuser() error {
	query := `
	INSERT INTO users (email, password) VALUES(?, ?)
	`
	stmt, err := db.DB.Prepare(query)

	defer stmt.Close()

	if err != nil {
		return err
	}

	hashedpassword, err := utils.HashPassword(u.Password)

	result, err := stmt.Exec(u.Email, hashedpassword)
	if err != nil {
		return err
	}

	_, err = result.LastInsertId()

	return err
}

func (u *User) Validateuserlogin() error {
	query := `
	SELECT id, password FROM users WHERE email = ?
	`
	row := db.DB.QueryRow(query, u.Email)
	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)

	if err != nil {
		return errors.New("credentails failed")
	}

	result := utils.Comparepassword(u.Password, retrivedPassword)

	if !result {
		return errors.New("credentails failed")
	}

	return nil

}
