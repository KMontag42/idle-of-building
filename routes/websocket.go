package routes

import (
	"encoding/json"
	"fmt"

	"github.com/kmontag42/idle-of-building/character"
	"github.com/kmontag42/idle-of-building/simulation"
	"github.com/kmontag42/idle-of-building/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func WebSockets(c echo.Context) error {
	websocket.Handler(func(ws *websocket.Conn) {
		defer ws.Close()
		for {
			// when we get the "start map" message, we need to run the map
			// and send the results back to the client
			msg := ""
			if err := websocket.Message.Receive(ws, &msg); err != nil {
				c.Logger().Error(err)
				break
			}
			var message utils.Message
			if err := json.Unmarshal([]byte(msg), &message); err != nil {
				c.Logger().Error(err)
				break
			}
			if message.Type == "start map" {
				character_xml := message.Data
				char, err := character.LoadCharacter(character_xml)
				if err != nil {
					c.Logger().Error(err)
					break
				}

				results := simulation.ExecuteMapForCharacter(
					&char,
					ws,
					c,
				)

				end_message := fmt.Sprintf(
					"Map completed. %f experience gained.",
					results.ExperienceGained,
				)

				utils.EmitMessage(ws, "battle-end", end_message)
				break
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
