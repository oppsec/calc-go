package main

import (
	"io/ioutil"
	"log"
	"fmt"
	"github.com/fatih/color"
	"github.com/inancgumus/screen"
)

func ascii() {

	// Get ascii.txt file content
	content, err := ioutil.ReadFile("ascii.txt")

	// Error if ascii.txt doens't exist
	if err != nil {
		log.Fatal(err)
	}
	
	text := string(content)
	slogan := "  Calculate fast with Go v1.0\n\n"

	screen.Clear()

	color.Cyan(text)
	color.Green(slogan)
	
}

func main() {

	ascii()

	fmt.Print("[-] Type your first number: ")
	var FirstNumber int
	fmt.Scan(&FirstNumber)

	fmt.Print("[-] Type your second number: ")
	var SecondNumber int
	fmt.Scan(&SecondNumber)

	add := FirstNumber + SecondNumber
	sub := FirstNumber - SecondNumber
	mul := FirstNumber * SecondNumber
	div := FirstNumber / SecondNumber

	fmt.Println("\n[-] Addition:", add, "\n[-] Subtration:", sub)
	fmt.Println("[-] Multiplication:", mul, "\n[-] Division:", div)
}