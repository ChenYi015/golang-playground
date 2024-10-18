package main

import "fmt"

func f(i int) {
	switch i {
	case 1:
		fmt.Println("i is 1")
	case 2:
		fmt.Println("i is 2")
	case 3:
		fmt.Println("i is 3")
	default:
		fmt.Println("i is not 1, 2, or 3")
	}
}

func g(i int) {
	switch i {
	case 1:
		fmt.Println("i is 1")
		fallthrough
	case 2:
		fmt.Println("i is 2 or i fell through from 1")
		fallthrough
	case 3:
		fmt.Println("i is 3 or fell through from above")
	default:
		fmt.Println("i is not 1, 2, or 3")
	}
}

func main() {
	f(3)
	fmt.Println("=======")
	g(3)
}
