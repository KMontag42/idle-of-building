package simulation

import (
	"fmt"
	"log"
	"time"

	"github.com/kmontag42/idle-of-building/types"
	"github.com/kmontag42/idle-of-building/utils"
	"golang.org/x/net/websocket"
)

func SimulateWave(
	hero *types.Character,
	enemies []types.Enemy,
	ws *websocket.Conn,
) (types.BattleResult, error) {
	heroWon := true
	for _, enemy := range enemies {
		if !Battle(hero, enemy) {
			log.Printf("%s has lost the battle\n", hero.Name)
			heroWon = false
			break
		}
	}

	time.Sleep(1 * time.Second)

	log.Printf("%s has cleared the wave\n", hero.Build.ClassName)
	if heroWon {
		err := utils.EmitMessage(
			ws,
			"battle",
			fmt.Sprintf("%s has cleared the wave", hero.Name),
		)
		if err != nil {
			log.Printf("error sending message: %v\n", err)
		}
	}

	return types.BattleResult{Character: *hero, Result: heroWon, Enemies: enemies}, nil
}
