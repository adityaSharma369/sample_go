package main

import "fmt"

func main() {

	fmt.Println("its jenkins world")
	//declaring a integer variable x
	var x int
	x = 3                //assigning x the value 3
	fmt.Println("x:", x) //prints 3

	//declaring a integer variable y with value 20 in a single statement and prints it
	var y int = 82
	fmt.Println("y:", y)

	//declaring a variable z with value 50 and prints it
	//Here type int is not explicitly mentioned
	var z = 51
	fmt.Println("z:", z)

	//Multiple variables are assigned in single line- i with an integer and j with a string
	var i, j = 68, "hello"
	fmt.Println("i and j:", i, j)
}
