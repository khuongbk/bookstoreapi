package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

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

	cus := getTotal(cusID, k, bookID)
	book := subtractBook(bookID, k)
	if cus.Id > 0 {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(cus); err != nil {
			panic(err)
		}
		return
	}
	if cus.Id == 0 && book.Id == 0 {
		// If we didn't find it, 404
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusNotFound)
		if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusNotFound, Text: "Plese check your Book ID or number of books you have ordered"}); err != nil {
			panic(err)
		}

	}
	if cus.Id == 0 && book.Id > 0 {
		http.Redirect(w, r, "http://www.google.com", 301)
		/*
			w.Header().Set("Content-Type", "application/json; charset=UTF-8")
			w.WriteHeader(http.StatusMovedPermanently)
			if err := json.NewEncoder(w).Encode(jsonErr{Code: http.StatusMovedPermanently, Text: "Plese register "}); err != nil {
				panic(err)
			}
		*/
	}
}
