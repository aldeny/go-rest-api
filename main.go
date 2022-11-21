package main

/* Aplikasi ini berdasarkan tutorial https://www.youtube.com/watch?v=dENoPS8aRL8&t=1173s */

import (
	"net/http"

	bookcontroller "github.com/aldeny/go-crud-library/controllers/book-controller"
	"github.com/aldeny/go-crud-library/models"
	"github.com/gorilla/mux"
)

func main() {
	models.ConnectDatabase()

	//Penggunaan MUX
	r := mux.NewRouter()

	//Router List
	r.HandleFunc("/books", bookcontroller.Index).Methods("GET")
	r.HandleFunc("/books/{id}", bookcontroller.Show).Methods("GET")
	r.HandleFunc("/books", bookcontroller.Create).Methods("POST")
	r.HandleFunc("/books/{id}", bookcontroller.Update).Methods("PUT")
	r.HandleFunc("/books/{id}", bookcontroller.Destroy).Methods("DELETE")

	http.ListenAndServe(":8000", r)
}
