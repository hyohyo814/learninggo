package main

import (
	"fmt"
	"math"
)

const s string = "constant"

func main() {
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println("d", d)

	fmt.Println("int64(d)", int64(d))

	fmt.Println("sin(d)", math.Sin(n))
}
