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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func bookShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var bookID int
	var err error
	if bookID, err = strconv.Atoi(vars["bookID"]); err != nil {
		panic(err)
	}
	book := findBook(bookID)
	if book.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(book); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}

func customerShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var cusID int
	var err error
	if cusID, err = strconv.Atoi(vars["cusID"]); err != nil {
		panic(err)
	}
	cus := findCus(cusID)
	if cus.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(cus); err != nil {
			panic(err)
		}
		return
	}

	// If we didn't find it, 404
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusNotFound)
	if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found"}); err != nil {
		panic(err)
	}

}
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
	bookID := createBook(book)
	book1 := findBook(bookID)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(book1); err != nil {
		panic(err)
	}
}

//
func customerCreater(w http.ResponseWriter, r *http.Request) {
	var cus Customer
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &cus); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}
	id := createCus(cus)
	cus1 := findCus(id)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(cus1); err != nil {
		panic(err)
	}
}

func bookBuy(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var cusID int
	var bookID int
	var k int
	var err error
	if cusID, err = strconv.Atoi(vars["cusID"]); err != nil {
		panic(err)
	}
	if bookID, err = strconv.Atoi(vars["bookID"]); err != nil {
		panic(err)
	}
	if k, err = strconv.Atoi(vars["mount"]); err != nil {
		panic(err)
	}

	cus := updateCustomer(cusID, k, bookID)
	book := updateBook(bookID, k)
	if book.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(findBook(cus.Id)); err != nil {
			panic(err)
		}
		return
	}
	if book.Id == 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Not Found CustomerId or Book"}); err != nil {
			panic(err)
		}
	}
}
