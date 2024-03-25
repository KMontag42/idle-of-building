package enemy

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"

	"github.com/kmontag42/idle-of-building/types"
)

func ReadMonsterData() []types.MonsterLevel {
	// read monster data from data/monsters-simple.csv
	file, err := os.Open("data/monsters-simple.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	// skip the header
	reader.Read()

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	monster_levels := []types.MonsterLevel{}
	for _, record := range records {
		level, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}
		damage, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Fatal(err)
		}
		evasion, err := strconv.Atoi(record[2])
		if err != nil {
			log.Fatal(err)
		}
		accuracy, err := strconv.Atoi(record[3])
		if err != nil {
			log.Fatal(err)
		}
		life, err := strconv.ParseFloat(record[4], 64)
		if err != nil {
			log.Fatal(err)
		}
		experience, err := strconv.Atoi(record[5])
		if err != nil {
			log.Fatal(err)
		}
		minion_life, err := strconv.Atoi(record[6])
		if err != nil {
			log.Fatal(err)
		}
		armor, err := strconv.Atoi(record[7])
		if err != nil {
			log.Fatal(err)
		}

		monster_levels = append(monster_levels, types.MonsterLevel{
			Level:      level,
			Damage:     damage,
			Evasion:    evasion,
			Accuracy:   accuracy,
			Life:       life,
			Experience: experience,
			MinionLife: minion_life,
			Armor:      armor,
		})
	}

	return monster_levels
}

func BuildEnemy(name string, level int, boss bool, monster_levels []types.MonsterLevel) types.Enemy {
	for _, monster_level := range monster_levels {
		if monster_level.Level == level {
			return types.Enemy{
				MonsterLevel: monster_level,
				Name:         name,
                                Boss: boss,
			}
		}
	}
	log.Fatalf("Monster level %d not found", level)
	panic("Monster level not found")
}
