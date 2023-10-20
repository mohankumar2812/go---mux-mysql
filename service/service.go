package service

import (
	"encoding/json"
	"fmt"
	"muxWithSql/config"
	"muxWithSql/utils"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	var user config.Users
	fmt.Println("requrest comming")
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rester := utils.BadRequest("Invalid Json")
		json.NewEncoder(w).Encode(rester)
		return
	}

	exist := config.IsExist(user.Mail)

	if exist {
		json.NewEncoder(w).Encode("user already exist")
		return
	}

	pass := utils.HashPassword(user.Passwd)
	user.Passwd = pass

	result, err := user.Create()

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func LoginUser(w http.ResponseWriter, r *http.Request) {

	var user config.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	backupUser := user

	exist := config.IsExist(user.Mail)

	if exist {

		userData, _ := user.GetUser()

		userExist := utils.CheckPwd(userData.Passwd, backupUser.Passwd)

		if userExist {
			loginData := &config.JWTToken{
				Email: user.Mail,
				RegisteredClaims: jwt.RegisteredClaims{
					ExpiresAt: jwt.NewNumericDate(time.Now().Add(100 * time.Minute)),
					Subject: string(user.Mail),
				},
			}

			token, err := utils.CreateToken(loginData)

			if err != nil {
				json.NewEncoder(w).Encode("Token not generated")
				return
			}

			json.NewEncoder(w).Encode(token)
			return
		} else {
			json.NewEncoder(w).Encode("Incorrect Password")
			return
		}

	} else {
		json.NewEncoder(w).Encode("user Not Exist")
		return
	}

}

func GetUsers(w http.ResponseWriter, r *http.Request) {

	result, err := config.GetAllUser()

	if err != nil {
		json.NewEncoder(w).Encode(err)
		fmt.Println(err)
		return
	}

	fmt.Println("users", result)

	json.NewEncoder(w).Encode(result)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var user config.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rester := utils.BadRequest("Invalid Json")
		json.NewEncoder(w).Encode(rester)
		return
	}

	JWT := r.Header["Authorization"]


	mail, err := utils.ValidateToken(string(JWT[0]))

	if mail != user.Mail{
		json.NewEncoder(w).Encode("invalid user")
		return
	}

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	result, err := user.GetUser()

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {

	var user config.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rester := utils.BadRequest("Invalid Json")
		json.NewEncoder(w).Encode(rester)
		return
	}

	JWT := r.Header["Authorization"]

	mail, err := utils.ValidateToken(string(JWT[0]))

	if mail != user.Mail{
		json.NewEncoder(w).Encode("invalid user")
		return
	}

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	result, err := user.UpdateUser()

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}

func DeleteUser(w http.ResponseWriter, r *http.Request) {

	var user config.Users

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		rester := utils.BadRequest("Invalid Json")
		json.NewEncoder(w).Encode(rester)
		return
	}

	result, err := user.Delete()

	if err != nil {
		json.NewEncoder(w).Encode(err)
		return
	}

	json.NewEncoder(w).Encode(result)

}
