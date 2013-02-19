package mcclane

import (
	// "encoding/json"
	"fmt"
	"github.com/ign/ipl-mcclane/brackets"
	"io/ioutil"
	"log"
	"net/http"
)

func RequestHandler(w http.ResponseWriter, r *http.Request) {
	SetCORHeaders(w, r)
	switch r.Method {
	case "GET":
		SetCORHeaders(w, r)
		log.Println("GET")
		data, err := FindBracket(r.URL.Path[len("/brackets/v6/api/"):])
		if err != nil {
			w.WriteHeader(404)
			log.Println(err)
			fmt.Fprintf(w, "404 Not found")
			return
		}
		w.WriteHeader(200)
		fmt.Fprintf(w, string(data))
	case "PUT":
		log.Println("PUT")
		if !checkAuth(r) {
			log.Println("Auth failed")
			w.WriteHeader(401)
			return
		}
		result, err := ReadBody(r)
		if err != nil {
			log.Println(err)
			return
		}
		UpdateBracket(result)
		out, _ := brackets.Format(result)
		w.WriteHeader(200)
		fmt.Fprintln(w, string(out))

	case "POST":
		log.Println("POST")
		if !checkAuth(r) {
			log.Println("Auth failed")
			w.WriteHeader(401)
			return
		}
		result, err := ReadBody(r)
		if err != nil {
			log.Println(err)
			return
		}
		InsertBracket(result)
		out, _ := brackets.Format(result)
		w.WriteHeader(200)
		fmt.Fprintln(w, string(out))
	case "DELETE":
		log.Println("DELETE")
		if !checkAuth(r) {
			log.Println("Auth failed")
			w.WriteHeader(401)
			return
		}
		SetCORHeaders(w, r)
		err := RemoveBracket(r.URL.Path[len("/brackets/v6/api/"):])
		if err != nil {
			log.Println("Error with Delete")
		}

	case "OPTIONS":
		SetCORHeaders(w, r)
		w.WriteHeader(200)
	}
}

func ReadBody(r *http.Request) (*brackets.Bracket, error) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	result, err := brackets.Parse(body)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return result, nil
}

func SetCORHeaders(w http.ResponseWriter, r *http.Request) {
	origin := r.Header.Get("Origin")
	w.Header().Set("Access-Control-Allow-Origin", origin)
	w.Header().Set("Access-Control-Allow-Methods", "PUT, GET, POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Credentials", "true")
	w.Header().Set("Content-type", "application/json")
}
