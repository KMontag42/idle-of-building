package routes

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/kmontag42/idle-of-building/simulation"
	"github.com/kmontag42/idle-of-building/types"
	"github.com/kmontag42/idle-of-building/utils"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

func handleStartMap(
	character_id int,
	map_type string,
	ws *websocket.Conn,
	c echo.Context,
) bool {
	map_info, err := simulation.GetMapInfo(map_type)
	if err != nil {
		c.Logger().Error(err)
		return false
	}
	char := types.CharactersMap[character_id]
	utils.EmitLife(ws, &char)
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
		err = utils.EmitExperience(ws, &char)
		if err != nil {
			c.Logger().Error(err)
			return false
		}
	}
	end_message := fmt.Sprintf(
		"Map %s. %f experience gained.",
		battle_result,
		results.ExperienceGained,
	)
	utils.EmitLife(ws, &char)
	utils.EmitMessage(ws, "battle-end", end_message)
	time.Sleep(5 * time.Second)
	return true
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
				map_success := handleStartMap(character_id, map_type, ws, c)
				for map_success {
					map_success = handleStartMap(character_id, map_type, ws, c)
				}
			}
		}
	}).ServeHTTP(c.Response(), c.Request())
	return nil
}
