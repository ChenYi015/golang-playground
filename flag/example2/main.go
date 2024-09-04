package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

var (
	b bool
	n int
	m int64
	u uint
	v uint64
	f float64
	s string
	t time.Duration
)

func main() {
	fmt.Println("args =", strings.Join(os.Args[1:], " "))

	flag.BoolVar(&b, "b", false, "help message for flag b")
	flag.IntVar(&n, "n", 1234, "help message for flag n")
	flag.Int64Var(&m, "m", 1234, "help message for flag m")
	flag.UintVar(&u, "u", 1234, "help message for flag u")
	flag.Uint64Var(&v, "v", 1234, "help message for flag v")
	flag.Float64Var(&f, "f", 3.14, "help message for flag f")
	flag.StringVar(&s, "s", "hello", "help message for flag s")
	flag.DurationVar(&t, "t", 1*time.Second, "help message for flag t")

	flag.Parse()
	fmt.Println("b =", b)
	fmt.Println("n =", n)
	fmt.Println("m =", m)
	fmt.Println("u =", u)
	fmt.Println("v =", v)
	fmt.Println("s =", s)
	fmt.Println("t =", t)
}
