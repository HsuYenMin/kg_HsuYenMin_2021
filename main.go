package main

import (
	"fmt"
	"os"
	"strings"
)

/* we don't have const map in golang */
var digitcharToPhone map[rune]string = map[rune]string{'0': "Zero", '1': "One", '2': "Two",
	'3': "Three", '4': "Four", '5': "Five", '6': "Six", '7': "Seven",
	'8': "Eight", '9': "Nine"}

/* convert a decimal string to a string of phones */
func toPhones(num string) string{
	ret := ""
	for _, digit := range num {
		/* error handling */
		if (digit < '0' || digit > '9') {
			fmt.Println("usage: scripts go run main.go [num...]")
			fmt.Println("[num...] should only contain decimal digits")
			os.Exit(0)
		}
		ret += digitcharToPhone[digit]
	}
	return ret
}

func main() {
	argsWithoutProg := os.Args[1:]
	if (len(argsWithoutProg) < 2) {
		fmt.Println("usage: scripts go run main.go [num...]\n")
		os.Exit(0)
	}
	var phonestring []string
	/* go through all command line arguments */
	for _, num := range argsWithoutProg {
		phonestring = append(phonestring, toPhones(num))
	}
	/* concate the result with ',' */
	fmt.Println(strings.Join(phonestring,","))
}
