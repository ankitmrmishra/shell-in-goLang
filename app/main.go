package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage
	scanner := bufio.NewScanner(os.Stdin)
	
	for {
		fmt.Print("$ ")
		
		if !scanner.Scan() {
			return
		}
		
		line := scanner.Text()
		words := strings.Fields(line)
		
		// Skip empty input
		if len(words) == 0 {
			continue
		}
		
		command := words[0]
		
		// Handle exit command
		if command == "exit" {
			break
		}
		
		// Handle echo command
		if command == "echo" {
			fmt.Println(strings.Join(words[1:], " "))
			continue
		}
		if(command == "type"){
			secondCommand := words[1]
			if(secondCommand == "echo" || secondCommand == "exit" || secondCommand == "type"){
				fmt.Println(secondCommand , "is a shell builtin")
			} else {
						fmt.Println(secondCommand, ": not found")	
			}
			continue
		}
		
		// Unknown command
		fmt.Println(command + ": command not found")
	}
}