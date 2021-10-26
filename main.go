package main

import (
	"bufio"
	"fmt"
	"myapp/doctor"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(doctor.Intro())
	for {
		userInput, _ := reader.ReadString('\n')
		// Windows
		userInput = strings.Replace(userInput, "\r\n", "", -1)
		// Linux
		userInput = strings.Replace(userInput, "\n", "", -1)
		response := doctor.Response(userInput)
		fmt.Println((response))
		if userInput == "quit" {
			break
		}
	}
}
