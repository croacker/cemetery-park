package api

import (
	"encoding/json"
	"log"
	"net/http"

	"croacker.com/cemetery-park/data"
)

// API -
type API struct {
	users    *data.UsersManager
	quarters *data.QuartersManager
	areas    *data.AreasManager
	burials  *data.BurialsManager
}

func NewAPI(db *data.DB) *API {
	users, err := data.NewUsersManager(db)
	handleError(err)
	quarters, err := data.NewQuartersManager(db)
	handleError(err)
	areas, err := data.NewAreasManager(db)
	handleError(err)
	burials, err := data.NewBurialsManager(db)
	handleError(err)

	return &API{
		users:    users,
		quarters: quarters,
		areas:    areas,
		burials:  burials,
	}
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

//Process error
func handleError(err error) {
	if err != nil {
		log.Println(err)
	}
}