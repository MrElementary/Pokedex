package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type 'exit' to quit")

	for {
		fmt.Print("Pokedex > ")

		scanner.Scan()
		input := scanner.Text()
		input = strings.ToLower(input)
		input_arr := strings.Fields(input)

		fmt.Printf("Your command was: %v\n", input_arr[0])

		if input_arr[0] == "exit" {
			break
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stdout, "Error:", err)
	}
}
