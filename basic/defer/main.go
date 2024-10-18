package main

import "fmt"

func main() {
	f()
}

func f() (s string) {
	defer func() {
		fmt.Println(s)
	}()

	return "Hello, World!"
}
