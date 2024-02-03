package server

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/raphhawk/famtree/internal/person/app"
	"github.com/raphhawk/famtree/internal/person/ports"
)

const Port = ":8080"

type Server struct {
	Router *echo.Echo
	Port   string
	Person ports.PersonService
}

func NewServer(port string) *Server {
	var person ports.PersonService
	person, err := app.New()
	if err != nil {
		log.Println(err)
	}
	return &Server{
		Router: echo.New(),
		Port:   port,
		Person: person,
	}
}

func (s *Server) Personel() ports.PersonService {
	return s.Person
}

func (s *Server) MapUrls() error {
	s.Router.GET(ports.GetPerson, s.GetPerson)
	return nil
}

func (s *Server) ListenAndServe() error {
	log.Println("Listening on port", s.Port)
	if err := s.MapUrls(); err != nil {
		return err
	}
	if err := http.ListenAndServe(s.Port, s.Router); err != nil {
		return err
	}
	return nil
}
