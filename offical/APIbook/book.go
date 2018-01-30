package main

type Book struct {
	Id              int     `json:"id"`
	Title           string  `json:"title"`
	Published_at    string  `json:"published_at"`
	Remain_quantity float64 `json:"remain_quantity"`
	Image_url       string  `json:"imge_url"`
	Price           float64 `json: "price"`
}

type Books []Book
