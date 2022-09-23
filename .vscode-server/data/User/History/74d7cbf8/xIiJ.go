package handler

import (
	"net/http"

	"github.com/gorilla/mux"

	jwtmiddleware "github.com/auth0/go-jwt-middleware"
    jwt "github.com/form3tech-oss/jwt-go"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST")
	router.Handle("/search", http.HandlerFunc(searchHandler)).Methods("GET")
	return router
}
