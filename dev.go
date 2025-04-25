package main

import (
	"fmt"
)

func main() {
	var s []string
	fmt.Println("uninit:", "!", s == nil, s, len(s) == 0)
	s = make([]string, 3)
	fmt.Println("uninit:", "!", s == nil, s, len(s) == 3)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	s = append(s, "d")
	s = append(s, "f")
	fmt.Println("apd:", s, "len:", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy", c)

	l := s[2:5]
	fmt.Println("s2:5", l, "len:", len(l))

	l = s[:5]
	fmt.Println("s:5", l, "len:", len(l))

	l = s[2:]
	fmt.Println("s2:", l, "len:", len(l))
	var a int = 10
	fmt.Printf("&a1=%p\n", &a)
	{
		a := 10
		fmt.Printf("&a2=%p\n", &a)
	}
	{
		a := 100
		fmt.Printf("======\n")
		fmt.Printf("b:%b\n", &a)
		fmt.Printf("c:%c\n", &a)
		fmt.Printf("d:%d\n", &a)
		fmt.Printf("o:%o\n", &a)
		fmt.Printf("x:%x\n", &a)
		fmt.Printf("X:%X\n", &a)
		fmt.Printf("U:%U\n", &a)
		fmt.Printf("s:%s\n", &a)
		fmt.Printf("t:%t\n", &a)
		fmt.Printf("v:%v\n", &a)
		fmt.Printf("T:%T\n", &a)
	}
}
