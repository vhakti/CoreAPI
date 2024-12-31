package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vhakti/CoreAPI/db"
	"github.com/vhakti/CoreAPI/models"
)

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	var users []models.User
	db.DB.Find(&users)
	json.NewEncoder(w).Encode(&users)

	//w.Write([]byte("get users"))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User //Just to infere the table from where to fetch the request
	params := mux.Vars(r)
	fmt.Println(params["id"])

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User does not exists"))
	} else {

		json.NewEncoder(w).Encode(&user)
	}
	//w.Write([]byte("get user"))
}

func PostUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	json.NewDecoder(r.Body).Decode(&user)
	createdUSer := db.DB.Create(&user)
	err := createdUSer.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		json.NewEncoder(w).Encode(&user)
	}

	//w.Write([]byte("Post"))

}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {
	var user models.User
	params := mux.Vars(r)

	db.DB.First(&user, params["id"])

	if user.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("User does not exists"))
		return
	}
	//db.DB.Unscoped().Delete(&user)  //force to delete it
	db.DB.Delete(&user) //logical deletion
	w.WriteHeader(http.StatusOK)

}
