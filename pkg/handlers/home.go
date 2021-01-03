package handlers

import (
	"fmt"
	"net/http"
)

// HelloWorld is traditional
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world!")

}
