package handlers

import (
	"fmt"
	"net/http"
)

// Person is person
func Person(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello person!")
}
