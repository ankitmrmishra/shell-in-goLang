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
		
		// Handle type command
		if command == "type" {
			// Check if we have enough arguments
			if len(words) < 2 {
				fmt.Println("type: missing argument")
				continue
			}
			
			targetCommand := words[1]
			
			// Check if it's a builtin command
			if isBuiltin(targetCommand) {
				fmt.Println(targetCommand, "is a shell builtin")
			} else {
				fmt.Println(targetCommand + ": not found")
			}
			continue
		}
		
		// Unknown command
		fmt.Println(command + ": command not found")
	}
}

// isBuiltin checks if a command is a shell builtin
func isBuiltin(cmd string) bool {
	builtins := []string{"echo", "exit", "type"}
	
	for _, builtin := range builtins {
		if cmd == builtin {
			return true
		}
	}
	
	return false
}