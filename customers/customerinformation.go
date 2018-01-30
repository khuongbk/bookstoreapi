package main

type Customer struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	Total     int    `json:"total"`
	Email     string `json:"email"`
	Book_list string `json:"book_list"`
}

type Book struct {
	Id              int    `json:"id"`
	Title           string `json:"title"`
	Published_at    string `json:"published_at"`
	Remain_quantity int    `json:"remain_quantity"`
	Image_url       string `json:"imge_url"`
	Price           int    `json: "price"`
	Author          string `json:"author"`
}
