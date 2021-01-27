package main

import (
	"github.com/fatih/color"
)

// SuccessMessage :: Prints a green message with [!]
func SuccessMessage(message string) {
	color.Green("[!] %v", message)
}

// ErrorMessage :: Prints a red message with [x]
func ErrorMessage(message string) {
	color.Red("[x] %v", message)
}