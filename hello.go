package main

import (
	"fmt"
	// "math"
	// "time"
)

const s string = "constant"

func main() {
	/*fmt.Println("1")
	fmt.Println(true&&false)
	fmt.Println(true||false)

	var b,c int=1,2
	fmt.Println(b,c)
	var e int
	fmt.Println(e)

	f:="apple"
	fmt.Println(f) */
	/*t := time.Now()
	switch  {
	case t.Hour()<12:
		fmt.Println("before")

	default:
	fmt.Println("after")
	}
	var a int =1
	fmt.Println(a)

	fmt.Println("okkkk")
	*/
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Println("default %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("bill")

}
