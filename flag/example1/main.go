package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Println("args =", strings.Join(os.Args[1:], " "))

	var b = flag.Bool("b", false, "help message for flag b")
	var n = flag.Int("n", 1234, "help message for flag n")
	var m = flag.Int64("m", 1234, "help message for flag m")
	var u = flag.Uint("u", 1234, "help message for flag u")
	var v = flag.Uint64("v", 1234, "help message for flag v")
	var f = flag.Float64("f", 3.14, "help message for flag f")
	var s = flag.String("s", "hello", "help message for flag s")
	var t = flag.Duration("t", 1*time.Second, "help message for flag t")

	flag.Parse()
	fmt.Println("b =", *b)
	fmt.Println("n =", *n)
	fmt.Println("m =", *m)
	fmt.Println("u =", *u)
	fmt.Println("v =", *v)
	fmt.Println("f =", *f)
	fmt.Println("s =", *s)
	fmt.Println("t =", *t)
}
