package main

import (
	"fmt"
	// "math"
	// "time"
)

const s string = "constant"

func main() {

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("bool")
		case int:
			fmt.Println("int")
		default:
			fmt.Printf("default %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI(s)

}
