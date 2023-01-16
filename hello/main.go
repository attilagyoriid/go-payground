package main

import (
	"fmt"
)

func main() {
	fmt.Println("Attila Gyori")
	printFirstNSquares(40)
	printPositiveSquaresLessThan(40)
	printAllPositiveDivisorsOf(2020)
	swapValuesByPointers()
	isHatNeeded()
}

func printFirstNSquares(n int) {

	for i := 1; i <= n; i++ {
		fmt.Println(i * i)
	}

}

func printPositiveSquaresLessThan(n int) {
	i := 1

	for i*i < 40 {

		fmt.Println(i * i)
		i++

	}
}

func printAllPositiveDivisorsOf(n int) {

	for i := 1; i <= n; i++ {
		if n%i == 0 {
			fmt.Println(i)
		}
	}
}

func swapValuesByPointers() {

	var a int = 1
	b := 2
	fmt.Printf("original: a: %d b: %d\n", a, b)
	pa := &a
	pb := &b
	a, b = *pb, *pa
	fmt.Printf("swapped by pointer: a: %d b: %d\n", a, b)

}

func isHatNeeded() {

	fmt.Println("Enter Teamparature (celsius): ")

	var celsius int

	fmt.Scanln(&celsius)
	if celsius < 10 {
		fmt.Printf("Wear Hat")
	} else {
		fmt.Printf("No Hat Needed")
	}

}

func printByteInGB(s uint64) {

}
