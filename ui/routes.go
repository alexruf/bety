package ui

import "net/http"

func (u WebUi) RegisterRoutes() {
	router := u.server.Router.PathPrefix("/ui").Subrouter()

	router.HandleFunc("/", u.handleGetRoot()).Methods(http.MethodGet)
}
