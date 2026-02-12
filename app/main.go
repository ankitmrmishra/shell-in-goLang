package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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
		words := parsefunction(line)

		// Skip empty input
		if len(words) == 0 {
			continue
		}

		command := words[0]

		switch command {

		// Handle exit command
		case "exit":
			return

		// Handle echo command
		case "echo":
			fmt.Println(strings.Join(words[1:], " "))
			continue

		// Handle type command
		case "type":
			handleType(words)
			continue
		case "pwd":
            handlePwd()
			continue

	    case "cd":
            changeDirectoray(words[1])
			continue
		}
	

		// Handle external commands
		if runExternal(command, words[1:]) {
			continue
		}

		// Unknown command
		fmt.Println(command + ": command not found")
	}
}

// parsing the single quotes here 
func parsefunction(line string) []string{
	var words []string
    var currentWord strings.Builder
	inSingleQuote := false
	inDoubleQuote := false

	for i := 0; i <len(line); i++{
		ch := line[i]
        if ch == '\\' && !inDoubleQuote && !inSingleQuote {
          i++  // Move to next character
       if i < len(line) {
           // Add the next character literally
           currentWord.WriteByte(line[i])
       }
       continue
		} else if ch == '\'' && !inDoubleQuote  {
           inSingleQuote = !inSingleQuote
		} else if ch == '"' && !inSingleQuote {
         inDoubleQuote = !inDoubleQuote
		} else if inSingleQuote || inDoubleQuote{
			currentWord.WriteByte(ch)
		} else if ch == ' ' || ch == '\t'{
			if currentWord.Len() > 0 {
				words = append(words, currentWord.String())
				currentWord.Reset()
			}
		} else {
			// Regular character outside quotes
			currentWord.WriteByte(ch)
		}
	}
	if currentWord.Len() > 0 {
		words = append(words, currentWord.String())
	}

	return words
}



// handleType processes the `type` builtin command
func handleType(words []string) {
	// Check if we have enough arguments
	if len(words) < 2 {
		fmt.Println("type: missing argument")
		return
	}

	targetCommand := words[1]
	path, err := exec.LookPath(targetCommand)

	// Check if it's a builtin command
	if isBuiltin(targetCommand) {
		fmt.Println(targetCommand, "is a shell builtin")
	} else if err == nil {
		fmt.Println(targetCommand + " is " + path)
	} else {
		fmt.Println(targetCommand + ": not found")
	}
}

// runExternal tries to execute an external program
func runExternal(command string, args []string) bool {
	path, err := exec.LookPath(command)
	if err != nil {
		return false
	}

	name := filepath.Base(path)

	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	_ = cmd.Run()
	return true
}

// isBuiltin checks if a command is a shell builtin
func isBuiltin(cmd string) bool {
	builtins := []string{"echo", "exit", "type", "pwd", "cd"}

	for _, builtin := range builtins {
		if cmd == builtin {
			return true
		}
	}

	return false
}


func handlePwd()  {
    path, err := os.Getwd()
    if err != nil {
        fmt.Println("pwd:", err)
        return
    }
    fmt.Println(path)
}

func changeDirectoray(dir string) {
	if dir == "~"{
        path, _ := os.UserHomeDir()
		os.Chdir(path)
		return
	}
	
	_, err := os.Stat(dir)
    if err != nil {
        fmt.Println("cd: " + dir + ": No such file or directory")
        return
    }
 os.Chdir(dir)

}
