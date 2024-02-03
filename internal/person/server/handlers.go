package server

import (
	"log"
	"strconv"

	"github.com/labstack/echo/v4"
	ae "github.com/raphhawk/famtree/cmd/common/error"
)

func (s *Server) GetPersons(ctx echo.Context) error   { return nil }
func (s *Server) GetPeople(ctx echo.Context) error    { return nil }
func (s *Server) SendPeople(ctx echo.Context) error   { return nil }
func (s *Server) AddPerson(ctx echo.Context) error    { return nil }
func (s *Server) UpdatePerson(ctx echo.Context) error { return nil }
func (s *Server) RemovePerson(ctx echo.Context) error { return nil }

func (s *Server) GetPerson(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		send := ae.InvalidPayload
		send.Error = err
		return ctx.JSONPretty(send.StatusCode, send, "\t")
	}
	res := s.Person.GetPersonById(id)
	return ctx.JSONPretty(res.Info.StatusCode, res, "\t")
}
