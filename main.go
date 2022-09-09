package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

func getInput(reader io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", nil
	}
	text := scanner.Text()
	if len(text) == 0 {
		return "", errors.New("Empty message is not allowed")
	}
	return text, nil
}

func printSay(message string) {
	initial := 0
	maxCount := 40
	final := maxCount
	parts := (len(message) / maxCount) + 1

	if final >= len(message) {
		final = len(message)
	}

	for i := 0; i <= final; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")

	for i := 0; i <= parts; i++ {
		if final > len(message) {
			final = len(message)
		}
		fmt.Printf("< %s >\n", message[initial:final])
		if final >= len(message) {
			break
		}
		initial = final
		final = initial + maxCount
	}
	for i := 0; i <= final; i++ {
		fmt.Printf("-")
	}
	fmt.Printf("\n")
}

func printFigure(animal string) {

	if animal == "tux" {

		fmt.Println("   \\")
		fmt.Println("    \\")
		fmt.Println("       .--.")
		fmt.Println("      |o_o |")
		fmt.Println("      |:_/ |")
		fmt.Println("     //   \\ \\")
		fmt.Println("    (|     | )")
		fmt.Println("   /'\\_   _/`\\")
		fmt.Println("   \\___)=(___/")
	} else {

		fmt.Println("    \\   ^__^")
		fmt.Println("     \\  (oo)_______")
		fmt.Println("        (__)\\     )\\/\\")
		fmt.Println("           ||----w |")
		fmt.Println("           ||     ||")
	}

}

func main() {
	animal := flag.String("animal", "cow", "Image of animal, to show")
	flag.Parse()

	message, err := getInput(os.Stdin, flag.Args()...)
	if err != nil {
		os.Exit(1)
	}

	if message == "" {
		fmt.Println("expect 'message' data")
		os.Exit(1)
	}
	printSay(message)
	printFigure(*animal)
}
