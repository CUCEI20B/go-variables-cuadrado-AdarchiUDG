package main

import "fmt"

func main() {
	var sideLength float64
	fmt.Scanln(&sideLength)
	fmt.Println(sideLength * sideLength)
}
