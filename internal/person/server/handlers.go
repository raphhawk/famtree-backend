package server

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	ae "github.com/raphhawk/famtree/cmd/common/error"
	"github.com/raphhawk/famtree/internal/person/ports"
)

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

func (s *Server) GetPeople(ctx echo.Context) error { return nil }

func (s *Server) SendPeople(ctx echo.Context) error {
	var people, result []ports.PersonDTO
	err := ctx.Bind(&people)
	if err != nil {
		log.Println(err)
		send := ae.InvalidPayload
		send.Error = err
		return ctx.JSONPretty(send.StatusCode, send, "\t")
	}
	for _, person := range people {
		res := s.Person.CreatePerson(person)
		result = append(result, res)
	}
	return ctx.JSONPretty(http.StatusOK, result, "\t")
}

func (s *Server) AddPerson(ctx echo.Context) error {
	var person ports.PersonDTO
	err := ctx.Bind(&person)
	if err != nil {
		log.Println(err)
		send := ae.InvalidPayload
		send.Error = err
		return ctx.JSONPretty(send.StatusCode, send, "\t")
	}
	res := s.Person.CreatePerson(person)
	return ctx.JSONPretty(res.Info.StatusCode, res, "\t")
}

func (s *Server) UpdatePerson(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		send := ae.InvalidPayload
		send.Error = err
		return ctx.JSONPretty(send.StatusCode, send, "\t")
	}
	var person ports.PersonDTO
	person.ID = id
	err = ctx.Bind(&person)
	if err != nil {
		log.Println(err)
		send := ae.InvalidPayload
		send.Error = err
		return ctx.JSONPretty(send.StatusCode, send, "\t")
	}
	var result []ports.PersonDTO
	if person.FirstName != "" && person.LastName != "" {
		result = append(result, s.Person.UpdatePersonName(person))
	}
	if person.Gender != "" {
		result = append(result, s.Person.UpdatePersonGender(person))
	}
	if person.Dob != "" {
		result = append(result, s.Person.UpdatePersonDob(person))
	}
	if person.Email != "" {
		result = append(result, s.Person.UpdatePersonEmail(person))
	}
	return ctx.JSONPretty(http.StatusOK, result, "\t")
}

func (s *Server) RemovePerson(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		log.Println(err)
		send := ae.InvalidPayload
		send.Error = err
		return ctx.JSONPretty(send.StatusCode, send, "\t")
	}
	res := s.Person.DeletePerson(id)
	return ctx.JSONPretty(res.Info.StatusCode, res, "\t")
}
