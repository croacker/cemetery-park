package main

import (
	"encoding/json"
	"log"
	"net/http"

	"croacker.com/cemetery-park/api"
	"croacker.com/cemetery-park/conf"
	"croacker.com/cemetery-park/data"
	"github.com/gorilla/mux"
)

type RespJ struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	configuration := conf.Get()
	log.Printf("Port '%s'", configuration.Port)
	log.Printf("DB path '%s'", configuration.Sqlite.DbPath)

	db := data.NewSqliteDB(configuration.Sqlite.DbPath)
	api := api.NewAPI(db)
	log.Printf("api '%s'", api)

	r := mux.NewRouter()
	r.HandleFunc("/", handler)
	port := ":" + configuration.Port
	log.Fatal(http.ListenAndServe(port, r))
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
