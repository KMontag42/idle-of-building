package simulation

import (
	"math/rand"
	"fmt"
	"log"

	"github.com/kmontag42/idle-of-building/character"
	"github.com/kmontag42/idle-of-building/enemy"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type MapInfo struct {
  Name string
  MinWaveCount int
  MaxWaveCount int
  WaveInfo enemy.WaveInfo
}

type MapResult struct {
	Results          []BattleResult
	ExperienceGained float64
	Victory          bool
}

func (mr MapResult) String() string {
	return fmt.Sprintf(
		"Results: %v\nExperienceGained: %f\nVictory: %t\n",
		mr.Results,
		mr.ExperienceGained,
		mr.Victory,
	)
}

var map1 MapInfo = MapInfo{
  Name: "Map 1",
  MinWaveCount: 1,
  MaxWaveCount: 5,
  WaveInfo: enemy.WaveInfo{
    MinWaveSize:  10,
    MaxWaveSize:  30,
    MinWaveLevel: 60,
    MaxWaveLevel: 80,
    Boss:         false,
  },
}

func ExecuteMapForCharacter(
	character *character.Character,
	ws *websocket.Conn,
	c echo.Context,
) MapResult {
	var map_waves []enemy.WaveInfo

        wave_count := rand.Intn(map1.MaxWaveCount - map1.MinWaveCount + 1) + map1.MinWaveCount
        log.Printf("wave count: %d\n", wave_count)
        map_waves = make([]enemy.WaveInfo, wave_count)
        for i := 0; i < wave_count; i++ {
          map_waves[i] = map1.WaveInfo
        }

        // add boss wave
        map_waves = append(map_waves, enemy.WaveInfo{
          MinWaveSize:  1,
          MaxWaveSize:  1,
          MinWaveLevel: 100,
          MaxWaveLevel: 100,
          Boss:         true,
        })

	var results []BattleResult

	for wave := range map_waves {
		enemies := enemy.CreateWave(map_waves[wave])
		result, err := SimulateWave(character, enemies, ws)
		if err != nil {
			log.Printf("error simulating wave: %v\n", err)
			break
		}
		results = append(results, result)
	}

	experience_gained := 0
	for _, result := range results {
		if result.Result {
			for _, enemy := range result.Enemies {
				experience_gained += enemy.Experience
			}
		}
	}

	victory := true
	for _, result := range results {
		if !result.Result {
			victory = false
			break
		}
	}

	return MapResult{Results: results, ExperienceGained: float64(experience_gained), Victory: victory}
}
