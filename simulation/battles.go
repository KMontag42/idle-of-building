package simulation

import (
	"github.com/kmontag42/idle-of-building/types"
)

func Battle(hero *types.Character, enemy types.Enemy) bool {
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
