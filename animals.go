package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}
type Cow struct {
	Name string
}

type Bird struct {
	Name string
}

type Snake struct {
	Name string
}

func (a Cow) Eat() {
	fmt.Print("grass")
}

func (a Bird) Eat() {
	fmt.Print("worms")
}

func (a Snake) Eat() {
	fmt.Print("mice")
}

func (a Cow) Move() {
	fmt.Print("walk")
}

func (a Bird) Move() {
	fmt.Print("fly")
}

func (a Snake) Move() {
	fmt.Print("slither")
}

func (a Cow) Speak() {
	fmt.Print("moo")
}

func (a Bird) Speak() {
	fmt.Print("peep")
}

func (a Snake) Speak() {
	fmt.Print("hsss")
}

func main() {

	animals := make([]Animal, 0)
	var command, name, typ string

	for {
		fmt.Printf(">")
		fmt.Scan(&command, &name, &typ)
		switch command {
		case "newanimal":
			switch typ {
			case "cow":
				cow := new(Cow)
				cow.Name = name
				animals = append(animals, cow)
			case "bird":
				bird := new(Bird)
				bird.Name = name
				animals = append(animals, bird)
			case "snake":
				snake := new(Snake)
				snake.Name = name
				animals = append(animals, snake)
			}
		case "query":
			var foundAnimal Animal
			for _, animal := range animals {
				if animal.Name == name {

				}
			}
		}

	}

}
