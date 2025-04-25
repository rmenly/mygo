package main

import (
	"fmt"
)

func main() {
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println("v1", v1)

	v2 := m["k2"]
	fmt.Println("v2", v2)

	v3 := m["k3"]
	fmt.Println("v3", v3)

	// clear(m)
	fmt.Println(m)

	_, prs := m["k1"]
	fmt.Println("prs:", prs)

	n := map[string]int{"foo1": 1, "bar": 2}
	fmt.Println("new map:", n)
}
