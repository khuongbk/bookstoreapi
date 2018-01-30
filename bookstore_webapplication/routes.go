package main

import "net/http"

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// routes for customers
type Routes []Route

var routes = Routes{

	Route{
		"CustomerShow",
		"GET",
		"/customers/{cusID}",
		customerShow,
	},
	Route{
		"BookBuy",
		"GET",
		"/customers/{cusID}/buy{bookID}:{mount}",
		bookBuy,
	},
	Route{
		"Index1",
		"GET",
		"/customers",
		Index1,
	},
	Route{
		"customerCreate",
		"POST",
		"/customers",
		customerCreater,
	},
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
}
