package simulation

import (
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
	}

	if hero.Life() <= 0 {
		return false
	} else {
		return true
	}
}
