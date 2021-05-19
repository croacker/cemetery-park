package routes

import (
	"log"
	"net/http"

	"croacker.com/cemetery-park/api"
	"croacker.com/cemetery-park/conf"
	"github.com/gorilla/mux"
)

type WithCORS struct {
	r *mux.Router
}

//NewCorsRoutes builds the CORS routes for the api
func NewCorsRoutes(api *api.API) *WithCORS {
	return &WithCORS{newRoutes(api)}
}

// NewRoutes builds the routes for the api
func newRoutes(api *api.API) *mux.Router {
	configuration := conf.Get()
	mux := mux.NewRouter()

	// client static files
	clientFolder := configuration.ClientFolder
	log.Printf("Clietn folder %s", clientFolder)
	mux.Handle("/", http.FileServer(http.Dir(clientFolder+"/"))).Methods("GET")
	mux.PathPrefix("/js").Handler(http.StripPrefix("/js/", http.FileServer(http.Dir(clientFolder+"/js/"))))
	mux.PathPrefix("/img").Handler(http.StripPrefix("/img/", http.FileServer(http.Dir(clientFolder+"/img/"))))
	mux.PathPrefix("/css").Handler(http.StripPrefix("/css/", http.FileServer(http.Dir(clientFolder+"/css/"))))
	mux.PathPrefix("/fonts").Handler(http.StripPrefix("/fonts/", http.FileServer(http.Dir(clientFolder+"/fonts/"))))
	mux.PathPrefix("/icons").Handler(http.StripPrefix("/icons/", http.FileServer(http.Dir(clientFolder+"/icons/"))))

	// api
	mux.HandleFunc("/api", ListApiHandler)
	apiRouter := mux.PathPrefix("/api").Subrouter()

	apiRouter.HandleFunc("/user", api.AllUsers).Methods("GET")
	apiRouter.HandleFunc("/user/{id}", api.UserByID).Methods("GET")
	apiRouter.HandleFunc("/user", api.CreateUser).Methods("POST")
	apiRouter.HandleFunc("/user/{id}", api.UpdateUser).Methods("PUT")
	apiRouter.HandleFunc("/user/{id}", api.DeleteUser).Methods("DELETE")

	apiRouter.HandleFunc("/quarter", api.AllQuarters).Methods("GET")
	apiRouter.HandleFunc("/quarter/{id}", api.QuarterByID).Methods("GET")
	apiRouter.HandleFunc("/quarter", api.CreateQuarter).Methods("POST")
	apiRouter.HandleFunc("/quarter/{id}", api.UpdateQuarter).Methods("PUT")
	apiRouter.HandleFunc("/quarter/{id}", api.DeleteQuarter).Methods("DELETE")

	apiRouter.HandleFunc("/area", api.AllAreas).Methods("GET")
	apiRouter.HandleFunc("/area/{id}", api.AreaByID).Methods("GET")
	apiRouter.HandleFunc("/area", api.CreateArea).Methods("POST")
	apiRouter.HandleFunc("/area/{id}", api.UpdateArea).Methods("PUT")
	apiRouter.HandleFunc("/area/{id}", api.DeleteArea).Methods("DELETE")

	apiRouter.HandleFunc("/burial", api.AllBurials).Methods("GET")
	apiRouter.HandleFunc("/burial/{id}", api.BurialByID).Methods("GET")
	apiRouter.HandleFunc("/burial", api.CreateBurial).Methods("POST")
	apiRouter.HandleFunc("/area/{id}", api.UpdateBurial).Methods("PUT")
	apiRouter.HandleFunc("/burial/{id}", api.DeleteBurial).Methods("DELETE")

	return mux
}

//YourHandler
func ListApiHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("API!\n"))
}

func (s *WithCORS) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	log.Printf("%s %s %s\n", req.RemoteAddr, req.Method, req.URL)
	if origin := req.Header.Get("Origin"); origin != "" {
		res.Header().Set("Access-Control-Allow-Origin", origin)
		res.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		res.Header().Set("Access-Control-Allow-Headers", "*")
	}

	if req.Method == "OPTIONS" {
		return
	}
	s.r.ServeHTTP(res, req)
}
