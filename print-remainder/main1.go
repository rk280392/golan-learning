package main

import "fmt"

func main() {
	for a:=10; a<=100; a++ {
		fmt.Printf("When %v is divided by 4, the remainder (aka modulus) is %v\n", a, a%4)
	}
}
