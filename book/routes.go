package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	Route{
		"BookCreate",
		"POST",
		"/books",
		bookCreater,
	},
	Route{
		"BookShow",
		"GET",
		"/books/{bookID}",
		bookShow,
	},
	// Route{
	// 	"Index",
	// 	"GET",
	// 	"/",
	// 	Index,
	// },
	Route{
		"bookCleaner",
		"DELETE",
		"/books/{bookID}",
		bookCleaner,
	},
	Route{
		"bookUpdater",
		"PUT",
		"/books/{bookID}",
		bookUpdater,
	},
}
