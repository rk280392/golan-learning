package main

import "fmt"

func main() {
	switch marks := 30; {
		case marks < 35:
			fmt.Println("fail")
		case marks > 35 && marks < 75:
			fmt.Println("good")
		case marks > 75:
			fmt.Println("excellent")
		}
	}

