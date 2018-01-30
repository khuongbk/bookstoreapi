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
		"PUT",
		"/books",
		bookCreater,
	},
	Route{
		"BookShow",
		"GET",
		"/books/{bookID}",
		bookShow,
	},
	Route{
		"CustomerShow",
		"GET",
		"/customers/{cusID}",
		customerShow,
	},
	Route{
		"BookBuy",
		"POST",
		"/customers/{cusID}/buy{bookID}:{mount}",
		bookBuy,
	},

	Route{
		"customerCreate",
		"PUST",
		"/customers",
		customerCreater,
	},
}
