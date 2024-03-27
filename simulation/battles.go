package simulation

import (
	"github.com/kmontag42/idle-of-building/types"
	"github.com/kmontag42/idle-of-building/utils"
	"golang.org/x/net/websocket"
)

func Battle(hero *types.Character, enemy types.Enemy, ws *websocket.Conn) bool {
	for hero.Life > 0 && enemy.Life > 0 {
		enemy.Life -= hero.Dps()
		hero.Life = hero.Life - enemy.Damage
		utils.EmitLife(ws, hero)
	}

	if hero.Life <= 0 {
		return false
	} else {
		return true
	}
}
