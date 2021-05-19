package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// AreaJSON
type AreaJSON struct {
	ParentId    uint   `json:"parentId"`
	Name        string `json:"name"`
	Description string `json:"password"`
	Coord       string `json:"coord"`
}

func (api *API) AllAreas(w http.ResponseWriter, req *http.Request) {
	entities := api.areas.AllAreas()
	toJSON(w, entities)
}

func (api *API) AreaByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	entity := api.areas.AreaByID(id)
	toJSON(w, entity)
}

func (api *API) CreateArea(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondata := AreaJSON{}
	err := decoder.Decode(&jsondata)
	handleError(err)
	entity := api.areas.CreateArea(jsondata.ParentId, jsondata.Name, jsondata.Description, jsondata.Coord)
	toJSON(w, entity)
}

func (api *API) UpdateArea(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	decoder := json.NewDecoder(req.Body)
	jsondata := AreaJSON{}
	err := decoder.Decode(&jsondata)
	handleError(err)
	entity := api.areas.UpdateArea(id, jsondata.ParentId, jsondata.Name, jsondata.Description, jsondata.Coord)
	toJSON(w, entity)
}

func (api *API) DeleteArea(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	entity := api.areas.DeleteArea(id)
	toJSON(w, entity)
}
