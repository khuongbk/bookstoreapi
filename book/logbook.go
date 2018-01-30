package main

import (
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"
)

/*
func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)

		log.Printf(
			"%s\t%s\t%s\t%s",
			r.Method,
			r.RequestURI,
			name,
			time.Since(start),
		)
	})
}
*/
type DataWeb struct {
	Method     string
	RequestURI string
	Name       string
	Time       time.Duration
}

var (
	IsDrop = true
)

func Logger(inner http.Handler, name string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		inner.ServeHTTP(w, r)
		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			panic(err)
		}
		defer session.Close()

		session.SetMode(mgo.Monotonic, true)
		if IsDrop {
			err = session.DB("dataweb").DropDatabase()
			if err != nil {
				panic(err)
			}
		}
		//
		c := session.DB("dataweb").C("data")
		index := mgo.Index{
			Key:        []string{"title"},
			Unique:     true,
			DropDups:   true,
			Background: true,
			Sparse:     true,
		}
		//
		err = c.EnsureIndex(index)
		if err != nil {
			panic(err)
		}
		//

		// Insert Datas
		err = c.Insert(&DataWeb{Method: r.Method, RequestURI: r.RequestURI, Name: name, Time: time.Since(start)})
		if err != nil {
			panic(err)
		}
		/*
			log.Printf(
				"%s\t%s\t%s\t%s",
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
			)
		*/
	})
}
