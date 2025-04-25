package main

import "fmt"

func getSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}
func main() {
	nextNum := getSeq()
	fmt.Println(nextNum())
	fmt.Println(nextNum())
	fmt.Println(nextNum())

	fmt.Println("===")
	nextNum2 := getSeq()
	fmt.Println(nextNum2())
	fmt.Println(nextNum2())

	fmt.Println("===")

	add := func(a, b int) int {
		return a + b
	}
	result := add(3, 5)
	fmt.Println(result)
	multiply := func(x, y int) int {
		return x * y
	}
	product := multiply(3, 5)
	fmt.Println(product)

	fmt.Println("--")
	calculate := func(op1 func(int, int) int, x, y int) int {
		return op1(x, y)
	}
	sum := calculate(add, 2, 8)
	fmt.Println(sum)

	sum = calculate(multiply, 2, 8)
	fmt.Println(sum)
	// 匿名函数
	diff1 := calculate(func(a, b int) int {
		return a - b
	}, 20, 5)
	fmt.Println(diff1)

	fmt.Println("--")
	var fib func(n int) int
	fib = func(n int) int {

		fmt.Println("*", n)
		if n < 2 {
			return n
		}
		a := fib(n - 1)
		fmt.Println("a", n, a)
		b := fib(n - 2)
		fmt.Println("#", n, a, b, "result:", a+b)
		return a + b
		// return fib(n-1) + fib(n-2)
	}
	// fmt.Println("fib:", fib(3))
	fmt.Println("fib:", fib(7))
}
