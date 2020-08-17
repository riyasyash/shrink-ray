package main

import (
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/riyasyash/shrink-ray/db"
	"github.com/riyasyash/shrink-ray/urlshortner"
)

func main() {
	r := mux.NewRouter()
	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})
	s := r.PathPrefix("/api").Subrouter()
	Db, err := db.GetDatabase()
	if err != nil {
		log.Fatal(err)
	}
	c := urlshortner.New(Db)

	s.HandleFunc("/shorten", c.Shorten).Methods("POST")

	r.PathPrefix("/").HandlerFunc(c.Redirect).Methods("GET")
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(header, methods, origins)(r)))

}
