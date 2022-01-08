## ex004_go_math using math/big as factorials gets big quickly


The code here is sampled partly from simple examples the go playground.

```
~/projects/learning-go/ex004_go_math $ cat factorial.go
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
```
I'm using math/big as factorials start small but quickly get big :
```
~/projects/learning-go/ex004_go_math $ go run .
Please enter number: 3
Got it, you entered :  3

6
~/projects/learning-go/ex004_go_math $ go run .
Please enter number: 4
Got it, you entered :  4

24
~/projects/learning-go/ex004_go_math $ go run .
Please enter number: 5
Got it, you entered :  5

120
~/projects/learning-go/ex004_go_math $ go run .
Please enter number: 51
Got it, you entered :  51

1551118753287382280224243016469303211063259720016986112000000000000
```
