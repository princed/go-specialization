package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (a Cow) Eat() {
	fmt.Println("grass")
}
func (a Cow) Move() {
	fmt.Println("walk")
}
func (a Cow) Speak() {
	fmt.Println("moo")
}

func (a Bird) Eat() {
	fmt.Println("worms")
}
func (a Bird) Move() {
	fmt.Println("fly")
}
func (a Bird) Speak() {
	fmt.Println("peep")
}

func (a Snake) Eat() {
	fmt.Println("mice")
}
func (a Snake) Move() {
	fmt.Println("slither")
}
func (a Snake) Speak() {
	fmt.Println("hsss")
}

func main() {
	animals := map[string]Animal{}
	for {
		command, name, param := prompt()
		var animal Animal

		switch command {
		case "newanimal":
			animal = createAnimal(param)
			if animal != nil {
				animals[name] = animal
				fmt.Println("Created it!")
			}
		case "query":
			animal = animals[name]
			if animal != nil {
				actAnimal(animal, param)
			} else {
				fmt.Printf(`No animal with name "%s"
`, name)
			}
		default:
			fmt.Printf(`Unknown command "%s"
`, command)
		}
	}
}

func actAnimal(animal Animal, request string) {
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

func createAnimal(animalType string) (animal Animal) {
	switch animalType {
	case "cow":
		animal = Cow{}
	case "bird":
		animal = Bird{}
	case "snake":
		animal = Snake{}
	default:
		fmt.Printf(`Unknown animal "%s"
`, animalType)
	}
	return
}

func prompt() (command, name, param string) {
	print("> ")
	fmt.Scan(&command, &name, &param)
	return
}
