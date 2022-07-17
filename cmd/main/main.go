package main

import (
	"log"
	"net/http"

	"github.com/MussaShaukenov/bookstore-api/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/junzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)

	// create a server
	log.Fatal(http.ListenAndServe("localhost:9010", r))

}
