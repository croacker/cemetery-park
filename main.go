package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type RespJ struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handler(w http.ResponseWriter, r *http.Request) {
	respJ := RespJ{
		Name: "name",
		Age:  10,
	}
	toJSON(w, respJ)
}

func toJSON(reposnseWriter http.ResponseWriter, v interface{}) {
	jsondata, err := json.Marshal(v)
	if err != nil {
		http.Error(reposnseWriter, err.Error(), http.StatusInternalServerError)
		return
	}
	reposnseWriter.Header().Set("Content-Type", "application/json")
	reposnseWriter.Write(jsondata)
}
