package routes

import "net/http"

func MainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Core API working with hot reloading"))

}
