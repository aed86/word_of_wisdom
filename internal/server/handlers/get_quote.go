package handlers

import (
	"fmt"
	"net/http"
)

type Handler struct {
}

func GetQuote(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
