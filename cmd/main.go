package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")

	fmt.Print("-> ")
	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	fmt.Println(text)

	if strings.Compare("upscale", text) == 0 {
		fmt.Println("So you want to upscale? That's fine by me :D")
	}

	if strings.Compare("downscale", text) == 0 {
		fmt.Println("Ok, I'll kill one of my agents D:")
	}
}
