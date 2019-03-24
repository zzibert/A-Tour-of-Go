package main

import "C"
import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

	// Create Animals
	Cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	Bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	Snake := Animal{food: "mice", locomotion: "slither", noise: "hsss"}

	var userQuery [2]string
	var problemFound bool

	for {

		userQuery, problemFound = getUserInput(">")

		if problemFound {
			continue
		}

		animalType := strings.ToLower(strings.TrimSpace(userQuery[0]))
		queryType := strings.ToLower(strings.TrimSpace(userQuery[1]))

		switch animalType {

		case "cow":
			processAnimal(Cow, queryType)
		case "bird":
			processAnimal(Bird, queryType)
		case "snake":
			processAnimal(Snake, queryType)
		default:
			fmt.Println("I don't know that animal")
			continue

		}

	}

}

func processAnimal(animal Animal, queryType string) {

	switch queryType {
	case "speak":
		fmt.Println(animal.Speak())
	case "eat":
		fmt.Println(animal.Eat())
	case "move":
		fmt.Println(animal.Move())
	default:
		fmt.Println("Unknown characteristic")
	}
}

func getUserInput(question string) ([2]string, bool) {

	problem := false
	var query [2]string

	scanner := bufio.NewScanner(os.Stdin)
	// Get user input .. using bufio.Scanner to get around issues with spaces
	fmt.Print(question)
	scanner.Scan()
	userInput := scanner.Text()

	userInputSplit := strings.Split(userInput, " ")

	// Check if there are enough items to return
	if len(userInputSplit) < 2 {
		problem = true
		fmt.Println("Error Found - not enough items found")
		return query, problem
	}

	query[0] = userInputSplit[0]
	query[1] = userInputSplit[1]

	return query, problem
}
