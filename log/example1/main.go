package main

import (
	"log"
	"os"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.Lmsgprefix)
	log.SetPrefix("Log Prefix ")
	log.SetOutput(os.Stdout)

	log.Printf("The output prefix for the standard logger is: [%v]\n", log.Prefix())
	log.Print("Hello, World!\n")
	log.Println("Hello, World!")
	log.Printf("Hello, %v!", "World")

	// log.Fatalln("Fatal!")

	// log.Panicln("Panic!")
}
