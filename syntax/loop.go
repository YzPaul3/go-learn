package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// 没有while
// for循环条件里不需要括号
// 条件里可以省略初始条件，结束条件，递增表达式
func toBinary(n int) string {
	result := ""
	for ; n >0; n/=2 {
		lsb := n % 2
		result = strconv.Itoa(lsb) + result
	}
	return result
}


// 省略初始条件，递增条件，相当于while
func printFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}

func forever() {
	// 相当于while(true)，死循环
	for {
		// do something
		fmt.Println(toBinary(13))
	}
}


func main() {
	fmt.Println(toBinary(13))
	printFile("aaa.txt")
}
