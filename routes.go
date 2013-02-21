package main

import (
	"flag"
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
	var portNumber int
	flag.IntVar(&portNumber, "port", 2121, "Default port is 2121")
	flag.Parse()
	mcclane.LoadConfig()

	fmt.Println("Welcome to the party, pal.")
	fmt.Println("Starting on port: ", portNumber)
	http.HandleFunc(LIST_PATH, mcclane.ListHandler)
	http.HandleFunc(API_PATH, mcclane.RequestHandler)
	http.HandleFunc(PING_PATH, system.Ping)
	err := http.ListenAndServe(fmt.Sprintf(":%d", portNumber), nil)
	if err != nil {
		fmt.Println("Could not start IPL-Mcclane on port", portNumber)
	}
}
