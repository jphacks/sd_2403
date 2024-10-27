package server

import (
	"context"
	"io"
	"sushi/service"
)

func (server *Server) NewService() []io.Closer {
	ctx, _ := context.WithCancel(context.Background())
	svc := service.NewService(server.db, server.log, server.config, ctx)
	// add all services that need to be closed
	toClose := []io.Closer{}
	server.service = svc
	return toClose
}
