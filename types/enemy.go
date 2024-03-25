package types

type MonsterLevel struct {
	Level      int
	Damage     float64
	Evasion    int
	Accuracy   int
	Life       float64
	Experience int
	MinionLife int
	Armor      int
}

type Enemy struct {
	MonsterLevel
	Name string
        Boss bool
}
