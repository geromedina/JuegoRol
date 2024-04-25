package model

type Player struct {
	Nickname   string
	Password   string
	UUID       string
	Email      string
	Characters []Character
}

type Character struct {
	Name   string
	HP     int
	Skills []Skill
	Breed  string
	Level  int
}

type Skill struct {
	Name        string
	Description string
	Button      string
	Value       int
	Healing     bool
}

type Enemy struct {
	Name        string
	Description string
	EnemySkill  Skill
	EnemyLevel  []Level
	HP          int
}

type Level struct {
	Name        string
	Description string
	Enemys      []Enemy
	Size        Size
	Number      int
	Checkpoints []Checkpoint
}

type Checkpoint struct {
	Number int
	Status bool
}

type Size struct {
	Height int
	Width  int
}
