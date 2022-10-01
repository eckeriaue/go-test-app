package main

import (
	"fmt"
	"math"
)

func getAllMathOperations(a, b int) (mathOperate map[string]int) {
	mathOperate = map[string]int{
		"sum": a + b,
		"dec": a - b,
		"mul": a * b,
		"div": a / b,
	}
	return
}

// func calculateSetValues (a,b int)

func returnalTestFunc(x, y int) func() {
	return func() {
		fmt.Print(x, y)
	}
}

func main() {

	returnalTestFunc(4, 3)()
	var testSlice map[string]int = getAllMathOperations(3, 4)

	fmt.Println(testSlice["dec"])
	fmt.Println("")
	fmt.Println("")

	money := map[string]int{
		"dollar": 1000,
		"eur":    1231,
		"rub":    12123,
	}

	_, status := money["rere"]

	fmt.Print(money["rub"], status)

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

	// if D > 0 {
	// 	var x1 float64
	// 	var x2 float64

	// 	x1 = (-b + math.Sqrt(D)) / (2 * a)
	// 	x2 = (-b - math.Sqrt(D)) / (2 * a)
	// 	fmt.Println("Ваше уравнение имеет два корня\nD = " + fmt.Sprint(D))
	// 	fmt.Println("X1:" + fmt.Sprint(x1))
	// 	fmt.Println("X2:" + fmt.Sprint(x2))
	// } else if D == 0 {
	// 	var x float64
	// 	x = (-b) / (2 * a)
	// 	fmt.Println("ваше уравнение имеет один корень")
	// 	fmt.Println("X: " + fmt.Sprint(x))
	// } else if D < 0 {
	// 	fmt.Println("ваше уравнение не имеет корней")
	// }

	switch {
	case D == 0:
		var x float64
		x = (-b) / (2 * a)
		fmt.Println("Ваше уравнение имеет один коренб")
		// fmt.Println("x: " + fmt.Sprint((x)))
		fmt.Printf("x: %f", x)
	case D > 0:
		var x1 float64
		var x2 float64
		x1 = (-b + math.Sqrt(D)) / (2 * a)
		x2 = (-b - math.Sqrt(D)) / (2 * a)
		// fmt.Println("Ваше уравнение имеет два корня\nD = " + fmt.Sprint(D))
		// fmt.Println("X1:" + fmt.Sprint(x1))
		// fmt.Println("X2:" + fmt.Sprint(x2))
		fmt.Printf("Ваше уравнение имеет два корня\nD = %f", D)
		fmt.Printf("X1: %f", x1)
		fmt.Printf("X2: %f", x2)
	case D < 0:
		fmt.Println("Уравнение не имеет корней")
	default:
		fmt.Println("Ошибка! Неверный тип данных")
	}

	fmt.Scan(&a)

}
