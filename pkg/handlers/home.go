package handlers

import (
	"fmt"
	"net/http"
)

// Home is home
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")
}
