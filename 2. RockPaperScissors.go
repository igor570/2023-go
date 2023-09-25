package main

import (
	"fmt"
	"math/rand"
)

func main() {
	playerChoice := playerTurn()
	fmt.Println("Player's choice:", playerChoice)
	computerChoice := computerTurn()
	fmt.Println("Computer's Choice: ", computerChoice)
	result := winDecider(playerChoice, computerChoice)
	fmt.Println(result)

}

func computerTurn() string {
	computerChoices := []string{"Rock", "Paper", "Scissors"}
	randomIndex := rand.Intn(len(computerChoices))
	return computerChoices[randomIndex]

}

func playerTurn() string {
	fmt.Printf("Enter your choice (Rock, Paper or Scissors): ")
	var choice string
	fmt.Scanln(&choice)
	if choice == "Rock" {
		return "Rock"
	} else if choice == "Paper" {
		return "Paper"
	} else if choice == "Scissors" {
		return "Scissors"
	} else {
		return "Error"
	}
}

func winDecider(playerChoice, computerChoice string) string {
	if playerChoice == "Paper" && computerChoice == "Rock" {
		return "Player wins"
	} else if playerChoice == "Paper" && computerChoice == "Scissors" {
		return "Computer wins"
	}
	if playerChoice == "Rock" && computerChoice == "Scissors" {
		return "Player wins"
	} else if playerChoice == "Rock" && computerChoice == "Paper" {
		return "Computer wins"
	}
	if playerChoice == "Scissors" && computerChoice == "Paper" {
		return "Player wins"
	} else if playerChoice == "Scissors" && computerChoice == "Rock" {
		return "Computer wins"
	}
	if playerChoice == computerChoice {
		return "Draw"
	}
	return "Error"
}
