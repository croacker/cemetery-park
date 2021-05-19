package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// QuarterJSON
type QuarterJSON struct {
	Name        string `json:"name"`
	Description string `json:"password"`
	Coord       string `json:"coord"`
}

func (api *API) AllQuarters(w http.ResponseWriter, req *http.Request) {
	entities := api.quarters.AllQuarters()
	toJSON(w, entities)
}

func (api *API) QuarterByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	entity := api.quarters.QuarterByID(id)
	toJSON(w, entity)
}

func (api *API) CreateQuarter(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondata := QuarterJSON{}
	err := decoder.Decode(&jsondata)
	handleError(err)
	entity := api.quarters.CreateQuarter(jsondata.Name, jsondata.Description, jsondata.Coord)
	toJSON(w, entity)
}

func (api *API) UpdateQuarter(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	decoder := json.NewDecoder(req.Body)
	jsondata := QuarterJSON{}
	err := decoder.Decode(&jsondata)
	handleError(err)
	entity := api.quarters.UpdateQuarter(id, jsondata.Name, jsondata.Description, jsondata.Coord)
	toJSON(w, entity)
}

func (api *API) DeleteQuarter(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	entity := api.quarters.DeleteQuarter(id)
	toJSON(w, entity)
}
