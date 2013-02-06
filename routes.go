package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/ign/ipl-mcclane/mcclane"
	"net/http"
)

const (
	INDEX_PATH = "/brackets/"
	V6_PREFIX  = "v6/"
	API_PATH = INDEX_PATH + V6_PREFIX + "api/"
	LIST_PATH = INDEX_PATH + V6_PREFIX + "api/list/"
)

func main() {
	fmt.Println("Welcome to the party, pal.")
	http.HandleFunc(LIST_PATH, mcclane.ListHandler)
	http.HandleFunc(API_PATH, mcclane.RequestHandler)
	http.ListenAndServe(":2121", nil)
}
