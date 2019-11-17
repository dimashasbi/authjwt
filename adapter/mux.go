package adapter

import (
	"AuthorizationJWT/engine"

	"github.com/gorilla/mux"

	"net/http"
)

// Handler using for make a Route
type (
	// Handler structure for Application Start Server
	Handler struct {
		Router   *mux.Router
		muxToken *token
	}
)

// InitializeServer Application
func (a *Handler) InitializeServer(f engine.EnginesFactory) {
	// add Engine
	a.muxToken = &token{f.NewTokenEngines()}
	a.Router = mux.NewRouter()
	a.SetURL()
}

// SetURL for reloading
func (a *Handler) SetURL() {
	a.POST("/gettoken", a.GetToken)
	a.POST("/checktoken", a.CheckToken)

}

// GET wraps the router for GET method
func (a *Handler) GET(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("GET")
}

// POST wraps the router for POST method
func (a *Handler) POST(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.Router.HandleFunc(path, f).Methods("POST")
}

// GetToken use for Get Token
func (a *Handler) GetToken(w http.ResponseWriter, r *http.Request) {
	a.muxToken.GetToken(w, r)
}

// CheckToken use for Check  Token
func (a *Handler) CheckToken(w http.ResponseWriter, r *http.Request) {
	a.muxToken.CheckToken(w, r)
}

// Run the app on it's router
func (a *Handler) Run(port string) {
	http.ListenAndServe(port, a.Router)
}
