package main

import (
	"fmt"
	"math"
)

func main() {

	findD()
	// var a int8 = 2
	// var b int64 = int64(a)
	// fmt.Print(b)
	// var first_string string
	// var second_string string

	// var c string
	// fmt.Scan(&first_string)
	// fmt.Scan(&second_string)

	// fmt.Scan(&c)
	// fmt.Println(fmt.Sprint(b) + fmt.Sprint(c))

	// if first_string == second_string {
	// 	fmt.Print("first_string is second string")
	// }

	// num := 4
	// if num > 3 {
	// 	fmt.Println("num is greater of 3")
	// } else {
	// 	fmt.Println("num is not greater of 3")
	// 	fmt.Println("first_string is not second_string")
	// }
}

func findD() {

	var a float64
	var b float64
	var c float64

	fmt.Println("set var of a:")
	fmt.Scan(&a)
	fmt.Println("set var of b:")
	fmt.Scan(&b)
	fmt.Println("set var of c:")
	fmt.Scan(&c)

	D := (b * b) - 4*(a*c)

	if D > 0 {
		var x1 float64
		var x2 float64

		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("Ваше уравнение имеет два корня\nD = " + fmt.Sprint(D))
		fmt.Println("X1:" + fmt.Sprint(x1))
		fmt.Println("X2:" + fmt.Sprint(x2))
	} else if D == 0 {
		var x float64
		x = (-b) / (2 * a)
		fmt.Println("ваше уравнение имеет один корень")
		fmt.Println("X: " + fmt.Sprint(x))
	} else if D < 0 {
		fmt.Println("ваше уравнение не имеет корней")
	}

}
