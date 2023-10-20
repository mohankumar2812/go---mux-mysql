package config

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Users struct {
	ID     int    `json:id`
	Name   string `json:name bson:"name,omitempty"`
	Mail   string `json:mail bson:"mail,omitempty"`
	Passwd string `json:passwd bson:"passwd,omitempty"`
	Phno   string `json:phno bson:"phno,omitempy"`
}

type JWTToken struct {
	Email string
	jwt.RegisteredClaims
}

func (user *Users) Create() (sql.Result, error) {

	row, err := DBInstance.Exec("INSERT INTO users(id,name,mail,passwd,phno) VALUES(?,?,?,?,?)", user.ID, user.Name, user.Mail, user.Passwd, user.Phno)

	if err != nil {
		fmt.Println("user creation error")
		return nil, err
	}

	return row, err
}

func GetAllUser() ([]Users, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()

	rows, err := DBInstance.QueryContext(ctx, "SELECT * FROM users")

	user := Users{}
	result := []Users{}

	for rows.Next() {

		var id int
		var name, mail, phno, passwd string

		err := rows.Scan(&id, &name, &mail, &passwd, &phno)

		if err != nil {
			return nil, err
		}

		user.ID = id
		user.Mail = mail
		user.Name = name
		user.Phno = phno

		result = append(result, user)
	}

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (user *Users) GetUser() (*Users, error) {

	rows := DBInstance.QueryRow("SELECT * FROM users WHERE mail=?", user.Mail)

	var id int
	var name, mail, phno, passwd string

	err := rows.Scan(&id, &name, &mail, &passwd, &phno)

	if err != nil {
		return nil, err
	}

	user.ID = id
	user.Mail = mail
	user.Name = name
	user.Passwd = passwd
	user.Phno = phno

	return user, nil

}

func (user *Users) UpdateUser() (*Users, error) {

	isExist := IsExist(user.Mail)

	if !isExist {
		return nil, fmt.Errorf("user not found")
	}

	prepareData, err := DBInstance.Prepare("UPDATE users SET name=?, mail=?, phno=? WHERE id=?")

	if err != nil {
		return nil, err
	}

	prepareData.Exec(user.Name, user.Mail, user.Phno, user.ID)

	updateUser, err := user.GetUser()

	if err != nil {
		return nil, err
	}

	return updateUser, nil

}

func (user *Users) Delete() (sql.Result, error) {

	prepareData, err := DBInstance.Prepare("DELETE FROM users WHERE mail=? ")

	if err != nil {
		return nil, err
	}

	result, err := prepareData.Exec(user.Mail)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func IsExist(mail string) bool {

	var user Users

	user.Mail = mail

	_, err := user.GetUser()

	return err == nil

}
