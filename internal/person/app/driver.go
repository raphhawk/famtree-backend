package app

import (
	"github.com/labstack/echo/v4"
	ae "github.com/raphhawk/famtree/cmd/common/error"
	"github.com/raphhawk/famtree/internal/person/ports"
)

func (p *Person) GetPerson(ctx echo.Context) error {
	var dto ports.PersonDTO
	err := ctx.Bind(&dto)
	if err != nil {
		send := ae.InvalidPayload
		send.Error = err
		return ctx.JSONPretty(send.StatusCode, send, "\t")
	}
	res := p.GetPersonById(dto.ID)
	return ctx.JSONPretty(200, res, "\t")
}
