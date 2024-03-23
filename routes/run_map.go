package routes

import (
	"net/http"

	"github.com/kmontag42/idle-of-building/character"
	"github.com/kmontag42/idle-of-building/simulation"
	"github.com/labstack/echo/v4"
)

func RunMapForCharacter(c echo.Context) error {
	character_xml := c.FormValue("xml")

	char, err := character.LoadCharacter(character_xml)
	if err != nil {
		return c.String(http.StatusNotAcceptable, err.Error())
	}
	results := simulation.ExecuteMapForCharacters([]character.Character{char}, nil, c)

	return c.Render(http.StatusOK, "map-results", results)
}
