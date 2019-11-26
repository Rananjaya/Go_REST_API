package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Book struct (Model) (Strct is Class in GO and like ES6 Class and same as classes in C#)
type Book struct {
	ID     string  `json:"id"`
	Isbn   string  `json:"isbn"`
	Title  string  `json:"title"`
	Author *Author `json:"author"` //*Author is own struct
}

//Auther struct

type Author struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

//Init Book var as a slice Book struct   (slice is veriyable like array)

var books []Book

//Get all Bookes  (any route hlder funcation shoud have request and response)
func getBook(w http.ResponseWriter, r *http.Request) {

}

func getsinglBook(w http.ResponseWriter, r *http.Request) {

}

func createBook(w http.ResponseWriter, r *http.Request) {

}

func updateBook(w http.ResponseWriter, r *http.Request) {

}
func deleteBook(w http.ResponseWriter, r *http.Request) {

}

func main() {

	//Init the router
	r := mux.NewRouter()

	//Mock Data -@todo - implement DB
	books = append(books, Book{ID: "1", Isbn: "44874", Title: "Book one", Author: &Author{FirstName: "Rananjaya", LastName: "Bandara"}})
	//create route hadelrs / End pointes
	r.HandleFunc("/api/bookes", getBook).Methods("GET")
	r.HandleFunc("/api/bookes/{id}", getsinglBook).Methods("GET") //get singel book
	r.HandleFunc("/api/bookes", createBook).Methods("POST")
	r.HandleFunc("/api/bookes/{id}", updateBook).Methods("PUT") //PUT use for update (ID sue for which one to update)
	r.HandleFunc("/api/bookes/{id}", deleteBook).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8080", r)) //Log.Fatal is used to throw an error
}
