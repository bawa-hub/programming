package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var pl = fmt.Println

func main() {
	// open, write and read a file
	f, err := os.Create("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	isPrimeArr := []int{2, 3, 5, 7}
	var sPrimeArr []string
	for _, i := range isPrimeArr {
		sPrimeArr = append(sPrimeArr, strconv.Itoa(i))
	}
	for _, num := range sPrimeArr {
		_, err := f.WriteString(num + "\n")
		if err != nil {
			log.Fatal(err)
		}
	}
	f, err = os.Open("data.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scan1 := bufio.NewScanner(f)
	for scan1.Scan() {
		pl("Prime: ", scan1.Text())
	}
	if err := scan1.Err(); err != nil {
		log.Fatal(err)
	}
}
