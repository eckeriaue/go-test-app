package main

import "fmt"

func main() {

	var a int8 = 2
	var b int64 = int64(a)
	fmt.Print(b)
	var first_string string
	var second_string string

	var c string
	fmt.Scan(&first_string)
	fmt.Scan(&second_string)

	fmt.Scan(&c)
	fmt.Println(fmt.Sprint(b) + fmt.Sprint(c))

	if first_string == second_string {
		fmt.Print("first_string is second string")
	}

	num := 4
	if num > 3 {
		fmt.Println("num is greater of 3")
	} else {
		fmt.Println("num is not greater of 3")
		fmt.Println("first_string is not second_string")
	}
}
