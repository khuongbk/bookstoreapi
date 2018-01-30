package main

type Customer struct {
	Id        int     `json:"id"`
	Name      string  `json:"name"`
	Total     float64 `json:"total"`
	Email     string  `json:"email"`
	Book_list string  `json:"book_list"`
}
type Customers []Customer
