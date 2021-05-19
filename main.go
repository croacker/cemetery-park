package main

import (
	"log"
	"net/http"

	"croacker.com/cemetery-park/api"
	"croacker.com/cemetery-park/conf"
	"croacker.com/cemetery-park/data"
	"croacker.com/cemetery-park/routes"
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
	router := routes.NewCorsRoutes(api)

	port := ":" + configuration.Port
	log.Fatal(http.ListenAndServe(port, router))
}
