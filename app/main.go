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
	line := scanner.Text()

	words := strings.Fields(line)
	if(len(words) == 0) {
		continue
	}

	if(words[0] == "exit"){
		break
	}
    if(words[0] == "echo"){
	fmt.Println(strings.Join(words[1:], " "))
			continue	
	}
	
	
	
	fmt.Println(words[0] + ": command not found")
	}
	
}
