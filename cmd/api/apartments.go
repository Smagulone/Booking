package main

import (
	"errors"
	"fmt"
	"github.com/Smagulone/Booking/internal/data"
	"github.com/Smagulone/Booking/internal/validator"
	"net/http"
)

func (app *application) bookApartmentHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "book an apartment")
}
func (app *application) showApartmentHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	apartment, err := app.models.Apartments.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"movie": apartment}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) updateMovieHandler(w http.ResponseWriter, r *http.Request) { // Extract the movie ID from the URL.
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Fetch the existing movie record from the database, sending a 404 Not Found // response to the client if we couldn't find a matching record.
	apartment, err := app.models.Apartments.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	// Copy the values from the request body to the appropriate fields of the movie // record.
	apartment.Name = input.Apartment
	movie.Year = input.Year
	movie.Runtime = input.Runtime
	movie.Genres = input.Genres
	// Validate the updated movie record, sending the client a 422 Unprocessable Entity // response if any checks fail.
	v := validator.New()
	if data.ValidateMovie(v, apartment); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	// Pass the updated movie record to our new Update() method.
	err = app.models.Movies.Update(apartment)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"apartment": apartment}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string   `json:"title"`
		Year    int32    `json:"year"`
		Runtime int32    `json:"runtime"`
		Genres  []string `json:"genres"`
	}
	err := app.readJSON(w, r, &input)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	apartment := &data.Apartment{
		Name:    input.Name,
		Ranking: input.Ranking,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}
	v := validator.New()
	if data.ValidateMovie(v, apartment); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}
	err = app.models.Apartment.Insert(apartment)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// When sending a HTTP response, we want to include a Location header to let the
	// client know which URL they can find the newly-created resource at. We make an
	// empty http.Header map and then use the Set() method to add a new Location header, // interpolating the system-generated ID for our new movie in the URL.
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/v1/movies/%d", apartment.ID))
	// Write a JSON response with a 201 Created status code, the movie data in the // response body, and the Location header.
	err = app.writeJSON(w, http.StatusCreated, envelope{"movie": apartment}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	err = app.models.Movies.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "apartment successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
