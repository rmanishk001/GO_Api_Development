package main

import (
	"api_dev/sub"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	fmt.Println("program is starting")
	sub.Test()

	r := mux.NewRouter()

	r.HandleFunc("/authdirectapi", login).Methods("Get")

	log.Fatal(http.ListenAndServe(":8000", r))

}
