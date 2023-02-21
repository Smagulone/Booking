package main

import (
	"fmt"
	"net/http"
)

func (app *application) bookHostelsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "book a hostel")
}
func (app *application) showHostelsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "show the details of hostels %d\n", id)
}
