package ui

import (
	"github.com/alexruf/bety/server"
)

type WebUi struct {
	server server.Server
}

func New(server server.Server) WebUi {
	return WebUi{
		server: server,
	}
}
