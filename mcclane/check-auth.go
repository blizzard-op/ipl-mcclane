package mcclane

import (
	"log"
	"net/http"
)

func checkAuth(r *http.Request) bool {
	authCookie, err := r.Cookie("robocopAuth")

	if err != nil {
		log.Println(err)
		return false
	}

	if authCookie.Value != "yourmovecreep" {
		return false
	}
	return true
}
