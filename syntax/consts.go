package main

import (
	"fmt"
	"math"
)

const typeC = "typeC"

func consts() {
	const filename = "aaa.txt"
	// const数值 可作为各种类型使用 int float
	const a, b = 3, 4
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c, filename)
	fmt.Println(typeC)
}

func enums() {
	const (
		cpp = iota //自增值,只能在常量的表达式中使用,在const关键字出现时将被重置为0
		java
		python
		golang
	)

	const (
		b = 1 << (10 * iota) // 自增值表达式
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, java, python, golang)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	consts()
	enums()
}
