package main

import "fmt"

const Pi = 3.141

func main() {
	var strHelloWorld = "Hello World"
	fmt.Println(strHelloWorld)

	strHelloWorld2 := "Hello World - with variable type inspection"
	fmt.Printf("Variable strHelloWorld2 is of type: %T\n", strHelloWorld2)

	var intMeaningOfLife = 42
	fmt.Println(intMeaningOfLife)
	fmt.Printf("Variable intMeaningOfLife is of type: %T\n", intMeaningOfLife)

	var intDefault int
	fmt.Println(intDefault)

	fmt.Println(Pi)

}
