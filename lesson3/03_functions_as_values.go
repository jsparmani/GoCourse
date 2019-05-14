package main

import (
	"fmt"
)

func testBinaryOps(a int, b int) {
	f1 := func(a int, b int) int { return a + b }
	f2 := func(a int, b int) int { return a - b }
	f3 := func(a int, b int) int { return a * b }
	f4 := func(a int, b int) int { return a / b }
	fmt.Printf("%d + %d = %d\n", a, b, f1(a, b))
	fmt.Printf("%d - %d = %d\n", a, b, f2(a, b))
	fmt.Printf("%d * %d = %d\n", a, b, f3(a, b))
	fmt.Printf("%d / %d = %d\n", a, b, f4(a, b))
}

func main() {
	testBinaryOps(1, 2)
}
