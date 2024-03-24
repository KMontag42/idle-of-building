package simulation

import (
	"log"

	"github.com/kmontag42/idle-of-building/character"
	"github.com/kmontag42/idle-of-building/enemy"
)

type BattleResult struct {
	Character character.Character
	Result    bool
	Enemies   []enemy.Enemy
}

func Battle(hero *character.Character, enemy enemy.Enemy) bool {
	for hero.Life() > 0 && enemy.Life > 0 {
		enemy.Life -= hero.Dps()
		hero.SetLife(hero.Life() - enemy.Damage)
		log.Printf("%s HP: %f\n", hero.Build.ClassName, hero.Life())
		log.Printf("%s HP: %f\n", enemy.Name, enemy.Life)
	}

	if hero.Life() <= 0 {
		log.Printf("%s has been defeated\n", hero.Build.ClassName)
		return false
	} else {
		log.Printf("%s has been defeated\n", enemy.Name)
		return true
	}
}
