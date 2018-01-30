package main

import (
	"fmt"
	"strings"
)

//create a type of book and insert the book into bookstock dbs
func createCustomer(c Customer) Customer {
	p, err := db.Prepare("INSERT customers set name=?,total=?,email=?,book_list=?")
	if err != nil {
		panic(err)
	}
	_, err = p.Exec(strings.Title(c.Name), c.Total, c.Email, c.Book_list)
	if err != nil {
		//panic(err)
		return Customer{}
	}
	id := countCustomers()
	cus := findCustomer(id)
	return cus
}

//find a Book
func findCustomer(cusID int) (cus Customer) {
	err := db.QueryRow("SELECT id,name,total,email,book_list FROM customers where id = ?", cusID).Scan(&cus.Id, &cus.Name, &cus.Total, &cus.Email, &cus.Book_list)
	if err != nil {
		//panic(err.Error()) // proper error handling instead of panic in your app
		return Customer{}
	}
	return cus
}

//
func findBook(bookID int) (book Book) {
	err := db.QueryRow("SELECT id, title,published_at,remain_quantity, image_url ,price ,author FROM bookstock where id = ?", bookID).Scan(&book.Id, &book.Title, &book.Published_at, &book.Remain_quantity, &book.Image_url, &book.Price, &book.Author)
	if err != nil {
		return Book{}
	}
	return book
}

//
func subtractBook(bookID int, mount int) Book {
	book := findBook(bookID)
	if book.Id > 0 {
		k := book.Remain_quantity - mount
		switch k > 0 {
		case true:
			book.Remain_quantity = k
			p, err := db.Prepare("update bookstock set remain_quantity=? where id=?")
			if err != nil {
				panic(err)
			}
			_, err = p.Exec(book.Remain_quantity, bookID)
			if err != nil {
				panic(err)
			}
			return book
			break
		case false:
			return Book{}
			break
		}
	}
	return Book{}
}

//
//subtrack books when customer buy "mount" books = update remain_quantity in bookstock where id=bookID
func getTotal(cusID int, mount int, bookID int) Customer {
	cus := findCustomer(cusID)
	book := subtractBook(bookID, mount)
	if cus.Id > 0 && book.Id > 0 {

		k := cus.Total + book.Price*mount
		t := strings.Contains(cus.Book_list, book.Title)
		switch t {
		case true:
			break
		case false:
			cus.Book_list = cus.Book_list + "  " + book.Title
			break
		}

		p, err := db.Prepare("update customers set total=?,book_list=? where id=?")
		if err != nil {
			panic(err)
		}
		_, err = p.Exec(k, cus.Book_list, cusID)
		if err != nil {
			panic(err)
		}

		cus1 := findCustomer(cusID)
		return cus1
	}
	return Customer{}
}
func updateCustomer(cusID int, newcus Customer) Customer {
	cus := findCustomer(cusID)
	switch cus.Id > 0 {
	case true:
		p, err := db.Prepare("update customers set name=?,email=? where id=?")
		if err != nil {
			panic(err)
		}
		_, err = p.Exec(newcus.Name, newcus.Email, cusID)
		if err != nil {
			panic(err)
		}
		cus = findCustomer(cusID)
		break
	case false:
		cus = Customer{}
		break
	}
	return cus
}

func countCustomers() (id int) {
	row := db.QueryRow("SELECT MAX(id) FROM customers ")
	err := row.Scan(&id)
	if err != nil {
		panic(err)
	}

	return id
}

//delete a book with bookID (delete a row in bookstock)
func deleteCustomer(cusID int) {
	p, err := db.Prepare("delete from customers where id=?")
	if err != nil {
		panic(err)
	}
	_, err = p.Exec(cusID)
	if err != nil {
		panic(err)
	}
}

type Data struct {
	Data []Customer `json:"data"`
}

func show() (Data, []Customer) {
	slice := make([]Customer, 0)
	var cus Customer
	var datacus Data
	id := countCustomers()
	for i := 1; i <= id; i++ {
		err := db.QueryRow("SELECT id,name,total,email,book_list FROM customers where id = ?", i).Scan(&cus.Id, &cus.Name, &cus.Total, &cus.Email, &cus.Book_list)
		if err != nil {
			//panic(err.Error()) // proper error handling instead of panic in your app
			fmt.Println("khuong")
		}

		slice = append(slice, cus)

	}
	datacus.Data = slice
	// datacus = {data:slice}
	return datacus, slice
	// return datacus
}
