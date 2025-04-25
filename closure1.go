package main

import "fmt"

func main() {
	fmt.Println("=======")
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println("=======")
	a := 1
	foo1(&a)()
	i := foo1(&a)
	fmt.Println(i)

}
func foo1(x *int) func() {
	return func() {
		*x = *x + 1
		fmt.Printf("foo1 val = %d\n", *x)
		fmt.Printf("foo1. val = \n", x)
		fmt.Printf("&x=%p\n", &x)
		fmt.Printf("...")
	}
}

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
