package main

import (
	"fmt"
	"github.com/ign/ipl-mcclane/mcclane"
	"github.com/ign/ipl-mcclane/system"
	"net/http"
)

const (
	INDEX_PATH = "/brackets/"
	V6_PREFIX  = "v6/"
	API_PATH   = INDEX_PATH + V6_PREFIX + "api/"
	LIST_PATH  = API_PATH + "list/"
	PING_PATH  = API_PATH + "ping/"
)

func main() {
	fmt.Println("Welcome to the party, pal.")
	http.HandleFunc(LIST_PATH, mcclane.ListHandler)
	http.HandleFunc(API_PATH, mcclane.RequestHandler)
	http.HandleFunc(PING_PATH, system.Ping)
	http.ListenAndServe(":2121", nil)
}
