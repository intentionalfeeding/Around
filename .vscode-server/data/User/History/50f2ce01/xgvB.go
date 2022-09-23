package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"around/model"
)

func uploadHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Received one post request")
	decoder := json.NewDecoder(r.body)
	var p model.Post
	if err := decoder.Decode(&p); err != nil{
	panic(err)

}

fmt.Fprintf(w,"Post received: %s\n", p.Message)
}


