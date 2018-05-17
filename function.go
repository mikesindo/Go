package main

import "fmt"

func add(a float32, b float32) float32 {
	return a + b
}

func main() {
	fmt.Println("sum is ", add(10.3, 19.3))
	fmt.Println(3 + 3)
}
