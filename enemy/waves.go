package enemy

import (
	"fmt"
	"math/rand"

	"github.com/kmontag42/idle-of-building/types"
)

func CreateWave(wave_info types.WaveInfo) (enemies []types.Enemy) {
	monster_levels := ReadMonsterData()
	number_of_enemies := max(int(rand.Float64()*float64(wave_info.MaxWaveSize)), wave_info.MinWaveSize)

	for i := 0; i < number_of_enemies; i++ {
		enemy_name := "Enemy " + fmt.Sprint(i)
		enemy_level := max(int(rand.Float64()*float64(wave_info.MaxWaveLevel)), wave_info.MinWaveLevel)
		enemy := BuildEnemy(enemy_name, enemy_level, wave_info.Boss, monster_levels)
		enemies = append(enemies, enemy)
	}

	return enemies
}
