package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("petedb> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\n", "", -1)
		firstIndex := string(text[0])

		if firstIndex == "." {
			HandleCommand(text)
		} else {
			message := fmt.Sprintf("Invalid statement: %v", text)
			fmt.Println(message)
		}
	}

}
