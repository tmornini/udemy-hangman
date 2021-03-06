package responses

import (
	"net/http"

	"github.com/tmornini/udemy-hangman/interfaces"
)

type Unauthorized struct {
	interfaces.Getable
}

// SerializeTo implement interfaces.Responsible
func (res Unauthorized) SerializeTo(w http.ResponseWriter) error {
	w.Header().Add("WWW-Authenticate", "Bearer realm=hangman")
	w.WriteHeader(401)

	return nil
}
