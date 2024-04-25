package main

import (
	"fmt"

	"github.com/geromedina/JuegoRol/model"
)

func main() {

}

func Menu() {
	var skills = []model.Skill{model.Skill{Name: "Corte brutal", Description: "lorem", Button: "a", Value: 20, Healing: false},
		model.Skill{Name: "Defensa enana", Description: "lorem", Button: "a", Value: 20, Healing: false},
		model.Skill{Name: "Rabia de las profundidades", Description: "lorem", Button: "a", Value: 20, Healing: false},
		model.Skill{Name: "Invocación de la luna", Description: "lorem", Button: "a", Value: 20, Healing: false}}
	var characters = []model.Character{model.Character{
		Name:   "Enano Luisito",
		HP:     150,
		Skills: []model.Skill{skills[0], skills[2]},
		Breed:  "Enano",
		Level:  10,
	}, model.Character{
		Name:   "Elfa Merlin",
		HP:     120,
		Skills: []model.Skill{skills[1], skills[3]},
		Breed:  "Elfa",
		Level:  8,
	}, model.Character{
		Name:   "Humano Javier",
		HP:     130,
		Skills: []model.Skill{skills[0], skills[3]},
		Breed:  "Humano",
		Level:  9,
	}, model.Character{
		Name:   "Mago Pedro",
		HP:     140,
		Skills: []model.Skill{skills[0], skills[2]},
		Breed:  "Mago",
		Level:  10,
	}}
	var nickname string
	fmt.Println("Bienvenido al juego")
	fmt.Println("Ingrese su nickname para continuar")
	fmt.Scanln(&nickname)
	fmt.Printf("Seleccione un personaje:\n1-")
}
