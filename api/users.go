package api

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// UserJSON
type UserJSON struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

func (api *API) AllUsers(w http.ResponseWriter, req *http.Request) {
	entities := api.users.AllUsers()
	toJSON(w, entities)
}

func (api *API) UserByID(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	entity := api.users.UserByID(id)
	toJSON(w, entity)
}

func (api *API) CreateUser(w http.ResponseWriter, req *http.Request) {
	decoder := json.NewDecoder(req.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)
	handleError(err)
	entity := api.users.CreateUser(jsondata.Name, jsondata.Password)
	toJSON(w, entity)
}

func (api *API) UpdateUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	decoder := json.NewDecoder(req.Body)
	jsondata := UserJSON{}
	err := decoder.Decode(&jsondata)
	handleError(err)
	entity := api.users.UpdateUser(id, jsondata.Name, jsondata.Password)
	toJSON(w, entity)
}

func (api *API) DeleteUser(w http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	id := params["id"]
	entity := api.users.DeleteUser(id)
	toJSON(w, entity)
}
