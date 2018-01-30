package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func bookShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookID int
	var err error
	if bookID, err = strconv.Atoi(vars["bookID"]); err != nil {
		fmt.Println("you had error:", err)
	}

	book := findBook(bookID)
	if book.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(book); err != nil {
			fmt.Println("you had error:", err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Plese check your Book's ID"}); err != nil {
		panic(err)
	}

}

//

func bookCreater(w http.ResponseWriter, r *http.Request) {
	var book Book
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &book); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	book1 := createBook(book)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(book1); err != nil {
		panic(err)
	}
	return

}

//

func bookCleaner(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookID int
	var err error
	if bookID, err = strconv.Atoi(vars["bookID"]); err != nil {
		panic(err)
	}
	book := findBook(bookID)
	if book.Id > 0 {
		deleteBook(bookID)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusOK, Text: "Deleted"}); err != nil {
			panic(err)
		}
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}
func bookUpdater(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookID int
	var err error
	if bookID, err = strconv.Atoi(vars["bookID"]); err != nil {
		panic(err)
	}
	var book Book
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &book); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	book1 := updateBook(bookID, book)
	if book1.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusCreated)
		if err := json.NewEncoder(w).Encode(book1); err != nil {
			panic(err)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}
