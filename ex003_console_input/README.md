## ex003_console_input -  `bufio` and `os` packages

Getting started with the `bufio` and `os` packages, a very simple example:
```
~/projects/learning-go/ex003_console_input $ cat bufio_os.go
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
```
works as expected
```
~/projects/learning-go/ex003_console_input $ go run .
Please enter command: testABC
Got it, you entered :  testABC

Let me get back to you on this one ...
```