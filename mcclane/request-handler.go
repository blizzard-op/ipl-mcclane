package mcclane

import (
	// "encoding/json"
	"github.com/ign/ipl-mcclane/brackets"
	"io/ioutil"
	"log"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	// case "GET":
	// 	result, err = getBracket(w, r, c)
	// case "PUT":
	// 	result, err = updateBracket(w, r, c)
	case "POST":
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
		}

		result, err := brackets.Parse(body)
		// gOut, _ := json.Marshal(result)
		// fmt.Println(string(gOut))
		if err != nil {
			log.Println(err)
		}
		// log.Println(result)
		CreateBracket(result)
		// FindBracket()
		SetCORHeaders(w)
		w.WriteHeader(200)
	case "OPTIONS":
		SetCORHeaders(w)
		w.WriteHeader(200)
	}
}

func SetCORHeaders(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-type", "application/json")
}
