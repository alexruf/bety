package server

import "net/http"

func (s *Server) RegisterRoutes() {
	router := s.Router.PathPrefix("/").Subrouter()

	router.HandleFunc("/", s.handleRedirect("/ui/")).Methods(http.MethodGet)

	router.HandleFunc("/health", s.handleGetHealth()).Methods(http.MethodGet)
}
