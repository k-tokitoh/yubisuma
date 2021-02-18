package main

import (
	"fmt"
)

func main() {
	fmt.Println("===== WELCOME TO YUBISUMA =====")
	for {
		game()
		if !askIfPlayAgain() {
			return
		}
	}
}

func game() {
	fmt.Println("game started.")
	var n int
	fmt.Scan(&n)
	fmt.Println(n)
	return
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
