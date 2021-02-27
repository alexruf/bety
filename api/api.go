package api

import (
	"github.com/alexruf/bety/server"
)

type Api struct {
	server server.Server
}

func New(server server.Server) Api {
	return Api{
		server: server,
	}
}
