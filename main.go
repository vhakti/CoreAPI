package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vhakti/CoreAPI/db"
	"github.com/vhakti/CoreAPI/routes"
)

func main() {

	db.DBConnection()
	r := mux.NewRouter() //Router
	//handle routes:
	r.HandleFunc("/", routes.MainHandler)

	http.ListenAndServe(":3000", r)
}
