package main

import (
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
	//routerbook := NewRouterforbook()

	//	log.Fatal(http.ListenAndServe(":8082", routerbook))
}
