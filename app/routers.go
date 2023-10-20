package app

import (
	"muxWithSql/service"
)

func Router() {

	router.HandleFunc("/register", service.CreateUser).Methods("POST")
	router.HandleFunc("/login", service.LoginUser).Methods("GET")
	router.HandleFunc("/getAllUsers", service.GetUsers).Methods("GET")
	router.HandleFunc("/getUser", service.GetUser).Methods("GET")
	router.HandleFunc("/updateUser", service.UpdateUser).Methods("PATCH")
	router.HandleFunc("/deleteUser", service.DeleteUser).Methods("DELETE")

}
