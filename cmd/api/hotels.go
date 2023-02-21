package main

import (
	"fmt"
	"net/http"
)

func (app *application) bookHotelsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "book a hotel")
}
func (app *application) showHotelsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of hotels %d\n", id)
}
