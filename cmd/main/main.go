package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rishabh-lt/go-bookstore/pkg/routes"
)

func main(){
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/",router)
	log.Fatal(http.ListenAndServe(":9010",router))
}