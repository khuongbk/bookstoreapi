package main

import (
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

/*
func init() {
	createBook(Book{Title: "harry porter", Remain_quantity: 50, Image_url: "https://www.golang-book.com/books/intro/7", Price: 50000})
	createCus(Customer{Name: "khuong", Email: "Nguyentienkhuong@gmail.com", Total: 0})
}
*/
func substractBook( bookID int , mount int){
 
}
func getTotal(cusID int , mount int, bookId)
func updateCustomer(cusID int, k int, bookID int) Customer {
	cus := findCus(cusID)
	book := findBook(bookID)
	cus.Total += book.Price * float64(k)
	if strings.Contains(cus.Book_list, book.Title) == false {
		cus.Book_list = cus.Book_list
	}
	cus.Book_list = cus.Book_list + "," + book.Title

	p, err := db.Prepare("update customers set total=?,book_list=? where id=?")
	if err != nil {
		panic(err)
	}
	_, err = p.Exec(cus.Total, cus.Book_list, cusID)
	if err != nil {
		panic(err)
	}
	cus = findCus(cusID)
	return cus
}
func updateBook(bookID int, n int)  {
	book := findBook(bookID)
	k := book.Remain_quantity - float64(n)
	switch k > 0 {
	case true:
	p, err := db.Prepare("update customers set remain_quantity=? where id=?")
	if err != nil {
		panic(err)
	}
	_, err = p.Exec(k, bookID)
	if err != nil {
		panic(err)
	}

		break

	case false:
			
			break
	}

	
}
func createBook(t Book) (id int) {
		p, err := db.Prepare("INSERT bookstock set title=?,price=?,published_at=?,remain_quantity=?,image_url=?")
		if err != nil {
			panic(err)
		}
		_, err = p.Exec(t.Title, t.Price, t.Published_at, t.Remain_quantity, t.Image_url)
		if err != nil {
			panic(err)
		}
		row := db.QueryRow("SELECT COUNT(*) FROM bookstock")
		err1 := row.Scan(&id)
		if err != nil {
			panic(err1)
			return 0
		}

		return id

}
func createCus(cus Customer) (id int) {

	p, err := db.Prepare("INSERT customers set name=?,email=?,total=?,book_list=?")
	if err != nil {
		panic(err)
	}
	_, err = p.Exec(cus.Name, cus.Email, cus.Total, cus.Book_list)
	if err != nil {
		panic(err)
	}
	row := db.QueryRow("SELECT COUNT(*) FROM customers")
	err1 := row.Scan(&id)
	if err != nil {
		panic(err1)
		return 0
	}

	return id
}

func deleteBook(bookId int) {
	stmt, err := db.Prepare("delete from bookstock where id=?")
	checkErr(err)

	res, err := stmt.Exec(bookId)
	checkErr(err)

	_, err = res.RowsAffected()
	checkErr(err)

	db.Close()
}

func findBook(id int) (book Book) {

	// Execute the query
	err := db.QueryRow("SELECT id, title,published_at,remain_quantity, image_url ,price FROM bookstock where id = ?", id).Scan(&book.Id, &book.Title, &book.Published_at, &book.Remain_quantity, &book.Image_url, &book.Price)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return Book{}
	}

	return book

}

func findCus(cusID int) (cus Customer) {

	// Execute the query
	err := db.QueryRow("SELECT id, name, total ,email,book_list FROM customers where id = ?", cusID).Scan(&cus.Id, &cus.Name, &cus.Total, &cus.Email, &cus.Book_list)
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
		return Customer{}
	}

	return cus

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
