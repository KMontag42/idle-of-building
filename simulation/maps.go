package simulation

import (
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"

	"github.com/kmontag42/idle-of-building/character"
	"github.com/kmontag42/idle-of-building/enemy"
)

type BattleResult struct {
	Character character.Character
	Result    bool
	Enemies   []enemy.Enemy
}

type MapResult struct {
	Results          []BattleResult
	ExperienceGained float64
	Victory          bool
}

func Battle(hero *character.Character, enemy enemy.Enemy) bool {
	for hero.Life() > 0 && enemy.Life > 0 {
		enemy.Life -= hero.Dps()
		hero.SetLife(hero.Life() - enemy.Damage)
		log.Printf("%s HP: %f\n", hero.Build.ClassName, hero.Life())
		log.Printf("%s HP: %f\n", enemy.Name, enemy.Life)
		time.Sleep(1 * time.Second)
	}

	if hero.Life() <= 0 {
		log.Printf("%s has been defeated\n", hero.Build.ClassName)
		return false
	} else {
		log.Printf("%s has been defeated\n", enemy.Name)
		return true
	}
}

func RunMap(
	hero *character.Character,
	enemies []enemy.Enemy,
	wg *sync.WaitGroup,
	resultChannel chan<- BattleResult,
) {
	defer wg.Done()
	heroWon := true
	for _, enemy := range enemies {
		log.Printf("%s has encountered %s\n", hero.Build.ClassName, enemy.Name)
		if !Battle(hero, enemy) {
			log.Printf("%s has lost the battle\n", hero.Build.ClassName)
			heroWon = false
			break
		}
	}
	log.Printf("%s has cleared the map\n", hero.Build.ClassName)
	resultChannel <- BattleResult{Character: *hero, Result: heroWon, Enemies: enemies}
}

func ExecuteMapForCharacters(characters []character.Character) MapResult {
	var wg sync.WaitGroup
	resultChannel := make(chan BattleResult)

	monster_levels := enemy.ReadMonsterData()
	enemies := []enemy.Enemy{}
	// generate random enemies
	number_of_enemies := 1 + int(rand.Float64()*19)
	for i := 0; i < number_of_enemies; i++ {
		enemy_name := "Enemy" + fmt.Sprint(i)
		enemy_level := int(rand.Float64() * 100.0)

		enemy := enemy.BuildEnemy(enemy_name, enemy_level, monster_levels)

		enemies = append(
			enemies,
			enemy,
		)
	}
	maps := map[*character.Character][]enemy.Enemy{}

	for _, character := range characters {
		maps[&character] = enemies
	}

	for hero, enemies := range maps {
		wg.Add(1)
		go RunMap(hero, enemies, &wg, resultChannel)
	}

	go func() {
		wg.Wait()
		close(resultChannel)
	}()

	results := []BattleResult{}
	for result := range resultChannel {
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
