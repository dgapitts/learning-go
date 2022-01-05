package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter command: ")
	input1, _ := reader.ReadString('\n')
	fmt.Println("Got it, you entered : ", input1)
	fmt.Println("Let me get back to you on this one ...")
}
