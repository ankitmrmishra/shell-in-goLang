package main

import (
	"fmt"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Print

func main() {
	// TODO: Uncomment the code below to pass the first stage

	for {
fmt.Print("$ ")
	var command string

	fmt.Scanln(&command)
	if(command == "exit"){
		break
	}
	if(command == "echo"){
			var echotext string
		fmt.Scanln(&echotext)
		fmt.Println(echotext)
		break
	}
	fmt.Println(command + ": command not found")
	}
	
}
