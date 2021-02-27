package api

import "net/http"

func (a Api) RegisterRoutes() {
	router := a.server.Router.PathPrefix("/api").Subrouter()

	router.HandleFunc("/content", a.handleApi()).Methods(http.MethodGet)
}
