package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/dustin/go-humanize"
)

func main() {
	fmt.Println("Attila Gyori")
	printFirstNSquares(40)
	printPositiveSquaresLessThan(40)
	printAllPositiveDivisorsOf(2020)
	swapValuesByPointers()
	// isHatNeeded()
	// printByteInGB()
	primitiveTypesArithmetic()
	walkingOnSunshine("I used to think maybe you loved me, now I know that it's true")
	walkingOnSunshine("Now, I don't want you back for the weekend, not back for a day")
	// intToBinary()

	reverseSlice([]int{1, 2, 3, 4, 5})
	reverseSliceInplace([]int{1, 2, 3, 4, 5})
	minMax()

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

func printByteInGB() {
	var s uint64
	fmt.Println("Enter available disk size in byte: ")
	fmt.Scanln(&s)
	fmt.Printf("Your available this size: %s.", humanize.Bytes(s))
}

// func printRealRootOf() {
// 	var a, b, c float64
// 	fmt.Println("Enter a: ")
// 	fmt.Scanln(&a)
// 	fmt.Println("Enter b: ")
// 	fmt.Scanln(&b)
// 	fmt.Println("Enter c: ")
// 	fmt.Scanln(&c)
// 	d := b*b - (4 * a * c)
// 	var r float64
// 	switch d {
// 	case d < 0:
// 		r = 0
// 	default:
// 		fmt.Println("It's after noon")
// 	}
// }

func primitiveTypesArithmetic() {

	var a float64 = 0

	for i := 0; i < 10; i++ {
		a += .1
	}

	fmt.Printf("%t\n", a == 1.0)
}

func walkingOnSunshine(s string) {

	fw, nl := firstWordsAndItsLenght(s)
	fmt.Printf("The first word is \"%s\", which consists of %d letter(s).\n", fw, nl)
	fmt.Printf("The second clause is: \"%s\"\n", secondCaluse(s))
	fmt.Printf("The last five characters of the string is: \"%s\"\n", lastFiveChar(s))

}

func firstWordsAndItsLenght(s string) (string, int) {
	var nl int
	for i := range s {

		if s[i] == ' ' {
			return s[0:i], nl
		}
		nl++
	}
	return s, nl
}

func secondCaluse(s string) string {

	return strings.Split(strings.Split(s, ",")[1], ",")[0]
}

func lastFiveChar(s string) string {

	return s[len(s)-5:]
}

func intToBinary() {
	var s uint16
	fmt.Println("Enter non-negative integer: ")
	fmt.Scanln(&s)

	fmt.Printf("Binary number: %016b\n", s)
	fmt.Print("Binary number: ")
	{

		for i := 1 << 15; i > 0; i = i / 2 {
			if int(s)&i != 0 {
				fmt.Print("1")
			} else {
				fmt.Print("0")
			}
		}
	}

}

func reverseSlice(is []int) {
	fmt.Printf("Before reverse %v: ", is)
	var reverse []int

	for i := len(is) - 1; i >= 0; i-- {
		reverse = append(reverse, is[i])
	}

	fmt.Printf("After reverse %v: ", reverse)
}

func reverseSliceInplace(is []int) {
	fmt.Printf("Before reverse: %v \n", is)

	for i, j := 0, len(is)-1; i < j; i, j = i+1, j-1 {
		is[i], is[j] = is[j], is[i]
	}

	fmt.Printf("After reverse: %v \n", is)
}

func minMax() {
	rand.Seed(time.Now().UnixNano())
	type sl [5]float32
	var matrix = make([]sl, 3)
	var min float32 = -100
	var max float32 = 100
	for i := 0; i < cap(matrix); i++ {
		for j := 0; j < cap(matrix[0]); j++ {
			matrix[i][j] = min + rand.Float32()*(max-min)
		}
	}
	fmt.Printf("matrix: %v \n", matrix)
	for i := 0; i < 3; i++ {
		fmt.Printf("%+.4f \n", matrix[i])
	}

	for i := 0; i < cap(matrix); i++ {
		for j := 0; j < cap(matrix[0]); j++ {
			if min < matrix[i][j] {
				min = matrix[i][j]
			}
			if max > matrix[i][j] {
				max = matrix[i][j]
			}
		}
	}
	fmt.Printf("min: %f max: %f \n", min, max)
}
