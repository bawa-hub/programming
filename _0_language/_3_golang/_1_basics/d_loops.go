package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println

func main() {

	// ----- FOR LOOPS -----
	// for initialization; condition; postStatement {BODY}

	// Print numbers 1 through 5
	for x := 1; x <= 5; x++ {
		pl(x)
	}
	// Do the opposite
	for x := 5; x >= 1; x-- {
		pl(x)
	}

	// For is used to create while loops as well
	fX := 0
	for fX < 5 {
		pl(fX)
		fX++
	}

	// While true loop (Infinite Loop) will be used for a guessing game
	reader := bufio.NewReader(os.Stdin)
	seedSecs := time.Now().Unix() // Returns seconds as int
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print("Guess a number between 0 and 50 : ")
		pl("Random Number is :", randNum)
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.TrimSpace(guess)
		iGuess, err := strconv.Atoi(guess)
		if err != nil {
			log.Fatal(err)
		}
		if iGuess > randNum {
			pl("Lower")
		} else if iGuess < randNum {
			pl("Higher")
		} else {
			pl("You Guessed it")
			break
		}

		// Cycle through an array with range
		// More on arrays later
		// We don't need the index so we ignore it
		// with the blank identifier _
		aNums := []int{1, 2, 3}
		for _, num := range aNums {
			pl(num)
		}
	}
}