package model

type Player struct {
	Nickname     string
	PlayerMallet Mallet
}

type Card struct {
	Name   string
	HP     int
	Attack int
}

type Mallet struct {
	Cards []Card
}
