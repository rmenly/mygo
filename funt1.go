package main

import "fmt"

func plus(a int, b int) int {
	return a + b

}
func plusPlus(a, b, c int) int {
	return a + b + c
}
func main() {
	res := plus(1, 2)
	fmt.Println(res)
	res = plusPlus(1, 2, 3)
	fmt.Println(res)

	a, b := vals(10)
	fmt.Println(a, b)

	_, c := vals(20)
	fmt.Println(c)
	fmt.Println("=======")
	sum(1, 2)
	sum(1, 2, 3)
	nums := []int{1, 2, 3, 5}
	sum(nums...)

}

func vals(a int) (int, int) {
	return a * 1, a * 10
}

func sum(nums ...int) {
	fmt.Print(nums, "  ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}
