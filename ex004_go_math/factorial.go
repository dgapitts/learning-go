package main

// https://go.dev/play/p/53TmmygltkR for factorial function
// https://go.dev/play/p/r5dG6X5YuF for ParseInt demo

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Please enter number: ")
	input1, _ := reader.ReadString('\n')
	fmt.Println("Got it, you entered : ", input1)
	int1, _ := strconv.ParseInt(strings.TrimSpace(input1), 0, 64)
	fmt.Println(factorial(big.NewInt(int1)))
}

func factorial(x *big.Int) *big.Int {
	n := big.NewInt(1)
	if x.Cmp(big.NewInt(0)) == 0 {
		return n
	}
	return n.Mul(x, factorial(n.Sub(x, n)))
}
