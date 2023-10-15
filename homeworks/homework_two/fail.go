package main

import "fmt"

func fail() {
	var a float64
	fmt.Scan(&a)
	switch {
	case a > 0:
		fmt.Println(a + 10)
	case a < 0:
		fmt.Println(a - 10)
	}
}
