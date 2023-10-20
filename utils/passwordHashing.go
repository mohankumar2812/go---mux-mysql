package utils

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) string {
	pwd := []byte(password)

	hash, err := bcrypt.GenerateFromPassword(pwd, 14)

	if err != nil {
		log.Fatal("error", err)
	}

	return string(hash)
}

func CheckPwd(hashPwd string, pwd string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashPwd), []byte(pwd))

	return err == nil
	
}
