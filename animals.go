package main

import (
	"fmt"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a Animal) Eat() string {
	return a.food
}

func (a Animal) Move() string {
	return a.locomotion
}

func (a Animal) Speak() string {
	return a.noise
}

func main() {

	var animal, action string

	cow := Animal{"grass", "walk", "moo"}
	bird := Animal{"worms", "fly", "peep"}
	snake := Animal{"mice", "slither", "hsss"}

	for {
		fmt.Scan(&animal, &action)
		switch animal {
		case "cow":
			switch action {
			case "speak":
				fmt.Println(cow.Speak())
			case "move":
				fmt.Println(cow.Move())
			case "eat":
				fmt.Println(cow.Eat())
			}
		case "bird":
			switch action {
			case "speak":
				fmt.Println(bird.Speak())
			case "move":
				fmt.Println(bird.Move())
			case "eat":
				fmt.Println(bird.Eat())
			}
		case "snake":
			switch action {
			case "speak":
				fmt.Println(snake.Speak())
			case "move":
				fmt.Println(snake.Move())
			case "eat":
				fmt.Println(snake.Eat())
			}
		}
	}

}
