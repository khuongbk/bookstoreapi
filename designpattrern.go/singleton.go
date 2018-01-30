package main

import "sync"

type book struct {
	state string
}

var singleton *book
var once sync.Once

func GetBook() *book {
	once.Do(func() {
		singleton = &book{state: "empty"}
	})
	return singleton
}

//get state book
func (sb *book) GetState() string {
	return sb.state
}

func (sb *book) SetState(s string) {
	sb.state = s
}

// no goliable variable here
func main() {

	sb := GetBook()
	if sb.GetState() == "empty" {
		println("book all sold")

	}
}
