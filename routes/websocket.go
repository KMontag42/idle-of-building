package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"strconv"

	"github.com/kmontag42/idle-of-building/simulation"
	"github.com/kmontag42/idle-of-building/types"
	"github.com/kmontag42/idle-of-building/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func EmitExperience(ws *websocket.Conn, character *types.Character) error {
	// build the "websocket-message" template with the message data
	var data bytes.Buffer

	template := template.Must(template.ParseFiles("views/pobpaste.html.tmpl"))
	err := template.ExecuteTemplate(&data, "experience-oob", character)
	if err != nil {
		return err
	}

	rendered := data.String()

	err = websocket.Message.Send(ws, rendered)
	if err != nil {
		return err
	}
	return nil
}

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
			var message types.Message
			if err := json.Unmarshal([]byte(msg), &message); err != nil {
				c.Logger().Error(err)
				break
			}
			if message.Type == "start map" {
				map_type := message.Map
                                character_id, err := strconv.Atoi(message.Id)
                                if err != nil {
                                  c.Logger().Error(err)
                                  break
                                }
				map_info, err := simulation.GetMapInfo(map_type)
				utils.EmitMessage(
					ws,
					"battle-start",
					fmt.Sprintf("Started %s", map_type),
				)
				if err != nil {
					c.Logger().Error(err)
					break
				}
                                char := types.CharactersMap[character_id]
				if err != nil {
					c.Logger().Error(err)
					break
				}

				results := simulation.ExecuteMapForCharacter(
					&char,
					map_info,
					ws,
					c,
				)

				battle_result := "completed"
				if !results.Victory {
					battle_result = "failed"
				} else {
					err = EmitExperience(ws, &char)
					if err != nil {
						c.Logger().Error(err)
						break
					}
				}

				end_message := fmt.Sprintf(
					"Map %s. %f experience gained.",
					battle_result,
					results.ExperienceGained,
				)

				utils.EmitMessage(ws, "battle-end", end_message)
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
