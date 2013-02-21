package mcclane

import (
	"log"
	"net/http"
)

func checkAuth(r *http.Request) bool {
	authCookie, err := r.Cookie(config.PassKey)

	if err != nil {
		log.Println(err)
		return false
	}

	if authCookie.Value != config.PassSecret {
		return false
	}
	return true
}
