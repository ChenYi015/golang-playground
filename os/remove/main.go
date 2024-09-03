package main

import (
	"fmt"
	"os"
)

func main() {
	err := os.Remove("nonexistentfile.txt")
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println("File does not exist.")
		} else {
			fmt.Println("Error:", err)
		}
	} else {
		fmt.Println("File removed successfully.")
	}
}
