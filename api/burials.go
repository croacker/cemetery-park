package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// BurialJSON
type BurialJSON struct {
	ParentId    uint   `json:"parentId"`
	Name        string `json:"name"`
	Description string `json:"password"`
	Coord       string `json:"coord"`
}

func (api *API) AllBurials(w http.ResponseWriter, req *http.Request) {
	entities := api.burials.AllBurials()
	toJSON(w, entities)
}

func (api *API) BurialByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	entity := api.burials.BurialByID(id)
	toJSON(w, entity)
}

func (api *API) CreateBurial(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondata := BurialJSON{}
	err := decoder.Decode(&jsondata)
	handleError(err)
	entity := api.burials.CreateBurial(jsondata.ParentId, jsondata.Name, jsondata.Description, jsondata.Coord)
	toJSON(w, entity)
}

func (api *API) UpdateBurial(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	decoder := json.NewDecoder(req.Body)
	jsondata := BurialJSON{}
	err := decoder.Decode(&jsondata)
	handleError(err)
	entity := api.burials.UpdateBurial(id, jsondata.ParentId, jsondata.Name, jsondata.Description, jsondata.Coord)
	toJSON(w, entity)
}

func (api *API) DeleteBurial(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	entity := api.burials.DeleteBurial(id)
	toJSON(w, entity)
}
