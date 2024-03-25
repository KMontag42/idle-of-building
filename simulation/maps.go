package simulation

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/kmontag42/idle-of-building/character"
	"github.com/kmontag42/idle-of-building/enemy"
	"github.com/labstack/echo/v4"
	"golang.org/x/net/websocket"
)

type MapInfo struct {
	Name         string
	MinWaveCount int
	MaxWaveCount int
	WaveInfo     enemy.WaveInfo
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

var WhiteMap MapInfo = MapInfo{
	Name:         "White Map",
	MinWaveCount: 1,
	MaxWaveCount: 5,
	WaveInfo: enemy.WaveInfo{
		MinWaveSize:  10,
		MaxWaveSize:  30,
		MinWaveLevel: 60,
		MaxWaveLevel: 70,
		Boss:         false,
	},
}

var YellowMap MapInfo = MapInfo{
	Name:         "Yellow Map",
	MinWaveCount: 4,
	MaxWaveCount: 8,
	WaveInfo: enemy.WaveInfo{
		MinWaveSize:  10,
		MaxWaveSize:  40,
		MinWaveLevel: 70,
		MaxWaveLevel: 80,
		Boss:         false,
	},
}

var RedMap MapInfo = MapInfo{
	Name:         "Red Map",
	MinWaveCount: 8,
	MaxWaveCount: 12,
	WaveInfo: enemy.WaveInfo{
		MinWaveSize:  10,
		MaxWaveSize:  50,
		MinWaveLevel: 80,
		MaxWaveLevel: 90,
		Boss:         false,
	},
}

var AllMaps []MapInfo = []MapInfo{WhiteMap, YellowMap, RedMap}

func GetMapInfo(name string) (MapInfo, error) {
  for _, map_info := range AllMaps {
    if map_info.Name == name {
      return map_info, nil
    }
  }
  return MapInfo{}, fmt.Errorf("map not found")
}

func ExecuteMapForCharacter(
	character *character.Character,
	map_info MapInfo,
	ws *websocket.Conn,
	c echo.Context,
) MapResult {
	var results []BattleResult

	// run a random number of waves
	wave_count := map_info.MinWaveCount + rand.Intn(
		map_info.MaxWaveCount-map_info.MinWaveCount,
	)

	for i := 0; i < wave_count; i++ {
		enemies := enemy.CreateWave(map_info.WaveInfo)
		result, err := SimulateWave(character, enemies, ws)
		if err != nil {
			log.Printf("error simulating wave: %v\n", err)
			break
		}
		results = append(results, result)

		// if the hero lost a battle, stop the simulation
		if !result.Result {
			break
		}
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
                        experience_gained = 0
			break
		}
	}

	return MapResult{Results: results, ExperienceGained: float64(experience_gained), Victory: victory}
}
