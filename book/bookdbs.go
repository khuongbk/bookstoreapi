package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:phamthuha27041104@tcp(localhost:3306)/data")
	if err != nil {
		panic(err)
	}
	//	createBook(Book{Title: "grammar in use", Remain_quantity: 50, Image_url: "https://www.golang-book.com/books/intro/7", Price: 50000, Published_at: "2016-12-10 ", Author: "Khuongg :v"})
	//	createCustomer(Customer{Name: "truong van luat", Email: "mabongdem@gmail.com"})
}

func Close() {
	db.Close()
}
