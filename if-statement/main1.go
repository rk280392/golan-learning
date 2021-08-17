package main

import "fmt"

func main() {
	x:=25
	if x < 2 {
		fmt.Println("smaller")
	} else if x > 23 {
		fmt.Println("Greater")
	} else {
		fmt.Println("equal")
	}
}
