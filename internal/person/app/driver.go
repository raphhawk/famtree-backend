package app

import (
	"strconv"

	"github.com/labstack/echo/v4"
	ae "github.com/raphhawk/famtree/cmd/common/error"
)

func (p *Person) GetPerson(ctx echo.Context) error {
	idStr := ctx.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		send := ae.InvalidPayload
		send.Error = err
		return ctx.JSONPretty(send.StatusCode, send, "\t")
	}
	res := p.GetPersonById(id)
	return ctx.JSONPretty(200, res, "\t")
}
