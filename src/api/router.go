package api

import (
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
)

const (
	staticPath = "./web"
	indexPath = "index.html"
)

// NewRouter will build a new router for the routes defined below
func NewRouter(c GameController) *mux.Router {

	r := mux.NewRouter()

	// Serve API routes
	api := r.PathPrefix("/v1/api/").Subrouter()
	api.HandleFunc("/game", c.NewGame).Methods("POST")
	api.HandleFunc("/game/{gameID}", c.GetGame).Methods("GET")
	//api.HandleFunc("/game/{gameID}", c.DeleteGame).Methods("DELETE")
	api.HandleFunc("/game/{gameID}/card", c.NewCard).Methods("POST")
	//api.HandleFunc("/game/{gameID}/card/{cardID}", c.UpdateCard).Methods("PUT")

	// Serve Frontend routes
	// For requests to dynamically generated game pages, serve index.html
	r.PathPrefix("/game/").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(staticPath, indexPath))
	})

	// Serve static build on root requests
	r.PathPrefix("/").Handler(http.FileServer(http.Dir(staticPath)))

	// TODO - Figure out how to serve styled 404 page for unhandled paths

	return r
}
