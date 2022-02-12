package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//login fuction
func login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var l Login

	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()

	err := decoder.Decode(&l)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(l.Message)

	var mainsh loginresponse
	mainsh.Username = "manish"
	mainsh.Message = "successfulyy"

	json.NewEncoder(w).Encode(&mainsh)

}
