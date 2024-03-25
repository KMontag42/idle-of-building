package types

import (
  "fmt"
)

type WaveInfo struct {
	MinWaveSize  int
	MaxWaveSize  int
	MinWaveLevel int
	MaxWaveLevel int
	Boss         bool
}
type MapInfo struct {
	Name         string
	MinWaveCount int
	MaxWaveCount int
	WaveInfo     WaveInfo
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
