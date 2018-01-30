package main

import (
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
)

// type bookfactory struct {

// }

//create a type of book and insert the book into bookstock dbs
func createBook(t Book) Book {
	p, err := db.Prepare("INSERT bookstock set title=?,price=?,published_at=?,remain_quantity=?,image_url=?,author=?")
	if err != nil {
		fmt.Println("you had error:", err)
	}
	_, err = p.Exec(strings.Title(t.Title), t.Price, t.Published_at, t.Remain_quantity, t.Image_url, t.Author)
	if err != nil {
		fmt.Println("you had error:", err)
	}
	id := countBooks()
	book := findBook(id)
	return book
}

//delete a book with bookID (delete a row in bookstock)
func deleteBook(bookID int) {
	p, err := db.Prepare("delete from bookstock where id=?")
	if err != nil {
		fmt.Println("you had error:", err)
	}
	_, err = p.Exec(bookID)
	if err != nil {
		fmt.Println("you had error:", err)
	}
}

//find a Book
func findBook(bookID int) (book Book) {
	err := db.QueryRow("SELECT id, title,published_at,remain_quantity, image_url ,price ,author FROM bookstock where id = ?", bookID).Scan(&book.Id, &book.Title, &book.Published_at, &book.Remain_quantity, &book.Image_url, &book.Price, &book.Author)
	if err != nil {
		return Book{}
	}
	return book
}

// subtrack books when customer buy "mount" books = update remain_quantity in bookstock where id=bookID
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
func updateBook(bookID int, newbook Book) Book {
	book := findBook(bookID)
	switch book.Id > 0 {
	case true:
		p, err := db.Prepare("update bookstock set remain_quantity=?, published_at=?,price=? where id=?")
		if err != nil {
			fmt.Println("you had error:", err)
		}
		_, err = p.Exec(newbook.Remain_quantity, newbook.Published_at, newbook.Price, bookID)
		if err != nil {
			fmt.Println("you had error:", err)
		}
		book = findBook(bookID)
		break
	case false:
		book = Book{}
		break
	}
	return book
}
func countBooks() (id int) {
	row := db.QueryRow("SELECT MAX(id) FROM bookstock ")
	err := row.Scan(&id)
	if err != nil {
		fmt.Println("you had error:", err)
	}

	return id
}
