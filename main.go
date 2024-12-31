package main

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vhakti/CoreAPI/db"
	"github.com/vhakti/CoreAPI/models"
	"github.com/vhakti/CoreAPI/routes"
)

func main() {

	db.DBConnection()
	db.DB.AutoMigrate(models.User{})
	db.DB.AutoMigrate(models.Tenant{})
	r := mux.NewRouter() //Router
	//handle routes:
	r.HandleFunc("/", routes.MainHandler)

	//users:
	r.HandleFunc("/users", routes.GetUsersHandler).Methods("GET")
	r.HandleFunc("/users/{id}", routes.GetUserHandler).Methods("GET")
	r.HandleFunc("/users", routes.PostUserHandler).Methods("POST")
	r.HandleFunc("/users/{id}", routes.DeleteUserHandler).Methods("DELETE")

	http.ListenAndServe(":3000", r)
}
