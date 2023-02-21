package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()
	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/apartments", app.bookApartmentHandler)
	router.HandlerFunc(http.MethodGet, "/v1/apartments/:id", app.showApartmentHandler)
	router.HandlerFunc(http.MethodPost, "/v1/hotels", app.bookHotelsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/hotels/:id", app.showHostelsHandler)
	router.HandlerFunc(http.MethodPost, "/v1/vila", app.bookVilaHandler)
	router.HandlerFunc(http.MethodGet, "/v1/vila/:id", app.showVilaHandler)
	router.HandlerFunc(http.MethodPost, "/v1/hostels", app.bookHostelsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/hostels/:id", app.showHostelsHandler)
	router.HandlerFunc(http.MethodPut, "/v1/apartments/:id", app.updateMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/movies/:id", app.deleteMovieHandler)
	return router
}
