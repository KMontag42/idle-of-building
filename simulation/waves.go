package simulation

import (
  "fmt"
  "log"
  "time"

  "github.com/kmontag42/idle-of-building/character"
  "github.com/kmontag42/idle-of-building/enemy"
  "github.com/kmontag42/idle-of-building/utils"
  "golang.org/x/net/websocket"
)

func SimulateWave(
	hero *character.Character,
	enemies []enemy.Enemy,
	ws *websocket.Conn,
) (BattleResult, error) {
	heroWon := true
	for _, enemy := range enemies {
		log.Printf("%s has encountered %s\n", hero.Build.ClassName, enemy.Name)
		if !Battle(hero, enemy) {
			log.Printf("%s has lost the battle\n", hero.Build.ClassName)
			heroWon = false
			break
		}
	}

	time.Sleep(1 * time.Second)

	log.Printf("%s has cleared the map\n", hero.Build.ClassName)
	if heroWon {
		err := utils.EmitMessage(
			ws,
			"battle",
			fmt.Sprintf("%s has defeated the wave", hero.Build.ClassName),
		)
		if err != nil {
			log.Printf("error sending message: %v\n", err)
		}
	}

	return BattleResult{Character: *hero, Result: heroWon, Enemies: enemies}, nil
}
