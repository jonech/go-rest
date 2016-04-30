package main

import (
	"fmt"
	"net/http"
)


// just for the welcome page
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

