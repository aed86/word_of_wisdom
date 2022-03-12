package handlers

import (
	"fmt"
	"net/http"
)

func GetChallenge(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
