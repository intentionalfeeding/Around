package handler

import (
    "net/http" 
    "gorilla/mux"   
)


func InitRouter() *mux.Router {
    router := mux.NewRouter()
    router.Handle("/upload", http.HandlerFunc(uploadHandler)).Methods("POST")
    return router
}
