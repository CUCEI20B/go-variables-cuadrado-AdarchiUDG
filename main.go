package main

import "fmt"

func main() {
	var sideLength uint32
	fmt.Scanln(&sideLength)
	fmt.Println(sideLength * sideLength)
}
