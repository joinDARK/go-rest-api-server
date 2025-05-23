package handlers

import (
	"fmt"
	"net/http"
)

func HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	}
}
