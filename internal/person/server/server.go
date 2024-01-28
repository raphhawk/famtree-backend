package server

import (
	"github.com/labstack/echo/v4"
	"github.com/raphhawk/famtree/internal/person/ports"
)

type Server struct {
	Router *echo.Echo
	Port   string
}

func NewServer(port string) *Server {
	return &Server{
		Router: echo.New(),
		Port:   port,
	}
}

func (s *Server) MapUrls() error {
	s.Router.GET(ports.GetPeople, GetPeople)
	return nil
}
