package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Data struct {
	Name string `json:"name"`
}

func RequestHandler(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	getData := Data{
		Name: "Aaryan Budhiraja",
	}
	data := Data{}
	switch req.Method {
	case "GET":
		err := json.NewEncoder(w).Encode(getData)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	case "POST":
		err := json.NewDecoder(req.Body).Decode(&data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
		fmt.Fprintf(w, "Input: %+v", data)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func main() {
	http.HandleFunc("/", RequestHandler)
	fmt.Println("Server Starting")
	http.ListenAndServe(":8080", nil)
}
