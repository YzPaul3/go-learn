package main

import "math"

// go没有隐式类型转换，只有强制类型转换，非常严格
func triangle() {
	var a, b int = 3, 4
	var c int
	c = int(math.Sqrt(float64(a * a + b * b)))
	println(c)
}

func main() {
	triangle()
}
