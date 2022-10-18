package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/LeonardoRafaelli/go-bookstore/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter();
	routes.RegisterBookStoreRoutes(r);
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}

