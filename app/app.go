package app

import (
	"log"
	"muxWithSql/config"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func App() {

	Router()
	config.DBconfig()
	log.Fatal(http.ListenAndServe(":6000", router))

}
