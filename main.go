package main

import (
	"github.com/geromedina/JuegoRol/model"
	"math/rand"
)

func main() {

}

func Menu() {

	var cards = []model.Card{
		model.Card{Name: "Tronador del abismo",
			HP:     50,
			Attack: 10},
		model.Card{Name: "Llamarada espectral",
			HP:     40,
			Attack: 15},
		model.Card{Name: "Guardian del trueno",
			HP:     68,
			Attack: 8},
		model.Card{Name: "Furia de las sombras",
			HP:     45,
			Attack: 20},
	}

}
func randomMallet(cards[]) Mallet {
	sizeCards := len(cards)
	var mallet = model.Mallet{Cards: }
	var cardsRandom = []model.Card
	for i := 0; i < 2; i++ {
		randomNumber := rand.Intn(sizeCards - 1)
		cardsRandom = append(cardsRandom, cards[randomNumber])

	}

}
