package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	welcome()
	for {
		game()
		if !askIfPlayAgain() {
			return
		}
	}
}

func welcome() {
	fmt.Println("=============================================")
	fmt.Println("==========   WELCOME TO YUBISUMA   ==========")
	fmt.Println("=============================================")
}

func game() {
	const playersCount = 3
	thumbsCounts := initGame(playersCount)

	for i := 0; i < 1000; i++ {
		var winnerIndex = winnerIndex(thumbsCounts)
		switch {
		case 0 <= winnerIndex:
			printResult(winnerIndex == 0)
			return
		default:
			time.Sleep(time.Millisecond * 1500)
			var indexOfplayerOnTurn = i % playersCount
			call(indexOfplayerOnTurn, thumbsCounts)
		}
	}
	return
}

func initGame(playersCount int) []int {
	fmt.Println(" -------------------------- ")
	fmt.Println("|     NEW GAME STARTED     |")
	fmt.Println(" -------------------------- ")

	thumbsCounts := make([]int, playersCount)
	for i := range thumbsCounts {
		thumbsCounts[i] = 2
	}
	return thumbsCounts
}

func printResult(win bool) {
	time.Sleep(time.Millisecond * 1500)

	switch win {
	case true:
		fmt.Println(" -------------------------- ")
		fmt.Println("|        YOU WIN!!!        |")
		fmt.Println(" -------------------------- ")
	case false:
		fmt.Println(" ------------------------- ")
		fmt.Println("|       you lose...       |")
		fmt.Println(" ------------------------- ")
	}
}

func winnerIndex(thumbsCounts []int) int {
	var winnerIndex int = -1
	for index, thumbsCount := range thumbsCounts {
		if thumbsCount == 0 {
			winnerIndex = index
			break
		}
	}
	return winnerIndex
}

func call(indexOfplayerOnTurn int, thumbsCounts []int) {
	isSelfTurn := isSelfTurn(indexOfplayerOnTurn)

	printTurnBeforehand(indexOfplayerOnTurn, thumbsCounts)

	upThumbsCounts := getUpThumbsCounts(thumbsCounts)

	guess := guessHowManyThumbsUp(isSelfTurn, sum(thumbsCounts))

	printCallResult(guess, thumbsCounts, upThumbsCounts)

	if isCallHit(guess, sum(upThumbsCounts)) {
		thumbsCounts[indexOfplayerOnTurn] -= 1
	}

	return
}

func isSelfTurn(indexOfplayerOnTurn int) bool {
	return indexOfplayerOnTurn == 0
}

func printTurnBeforehand(indexOfplayerOnTurn int, thumbsCounts []int) {
	fmt.Print("\n\n***************************\n\n")
	printTurn(thumbsCounts, make([]int, len(thumbsCounts)), indexOfplayerOnTurn)
}

func printTurn(thumbsCounts []int, upThumbsCounts []int, indexOfPlayerOnTurn int) {
	for i, thumbsCount := range thumbsCounts {
		var playerName string
		var turnMessage string

		if i == 0 {
			playerName = "you"
			turnMessage = "YOUR TURN!"
		} else {
			playerName = fmt.Sprintf("opponent %d", i)
			turnMessage = "his/her turn."
		}
		thumbs := strings.Repeat("UP   ", upThumbsCounts[i]) + strings.Repeat("__   ", thumbsCount-upThumbsCounts[i])
		description := fmt.Sprintf("%-13s%s", playerName, thumbs)
		if i == indexOfPlayerOnTurn {
			description += fmt.Sprintf("<= %s", turnMessage)
		}
		fmt.Printf("%s\n\n", description)
	}
}

func sum(arr []int) int {
	sum := 0
	for _, n := range arr {
		sum += n
	}
	return sum
}

func getUpThumbsCounts(thumbsCounts []int) []int {
	upThumbsCounts := []int{askHowManySelfThumbsUp(thumbsCounts[0])}
	for i := 1; i < len(thumbsCounts); i++ {
		upThumbsCounts = append(upThumbsCounts, rand.Intn(thumbsCounts[i]))
	}
	return upThumbsCounts
}

func askHowManySelfThumbsUp(max int) int {
	fmt.Printf("\nhow many thumbs will you go up? (enter 0 - %d)\n", max)
	var upSelfThumbsCount int
	fmt.Scan(&upSelfThumbsCount)
	fmt.Print("\n\n")
	return upSelfThumbsCount
}

func guessHowManyThumbsUp(isSelfTurn bool, thumbsSum int) int {
	var guess int
	if isSelfTurn {
		fmt.Printf("\nhow many thumbs do you guess will go up in all? (enter x - x)\n")
		fmt.Scan(&guess)
		fmt.Print("\n\n")
	} else {
		guess = rand.Intn(thumbsSum)
	}
	return guess
}

func printCallResult(guess int, thumbsCounts []int, upThumbsCounts []int) {
	var words = []string{"I ", "DO ", "GUESS ", fmt.Sprintf("<<< %d >>> !!!\n\n", guess)}
	var message string
	if isCallHit(guess, sum(upThumbsCounts)) {
		message = "HIT"
	} else {
		message = "MISS"
	}

	for _, word := range words {
		time.Sleep(time.Millisecond * 500)
		fmt.Print(word)
	}
	fmt.Print("\n\n")

	printTurn(thumbsCounts, upThumbsCounts, -1)

	time.Sleep(time.Millisecond * 2000)

	fmt.Printf("\n\n[[  %s!!  ]]\n\n", message)
}

func isCallHit(guess int, upThumbsSum int) bool {
	return upThumbsSum == guess
}

func askIfPlayAgain() bool {
	for {
		fmt.Println("play again? (enter y/n)")
		var c string
		fmt.Scan(&c)
		if c == "y" {
			return true
		} else if c == "n" {
			return false
		}
	}
}
