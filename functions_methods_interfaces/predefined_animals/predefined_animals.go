package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, noise string
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}
func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}
func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		animalType, request := prompt()
		var animal Animal

		switch animalType {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			fmt.Printf(`Unknown animal "%s"
`, animalType)
			continue
		}

		switch request {
		case "eat":
			animal.Eat()
		case "move":
			animal.Move()
		case "speak":
			animal.Speak()
		default:
			fmt.Printf(`Unknown request "%s"
`, request)
		}
	}
}

func prompt() (animal, request string) {
	print("> ")
	fmt.Scan(&animal, &request)
	return
}
