package main

import "fmt"

func main() {
	var a int8 = 2
	var b int64 = int64(a)
	fmt.Print(b)

	var c string

	fmt.Scan(&c)
	fmt.Println(fmt.Sprint(b) + fmt.Sprint(c))

}
