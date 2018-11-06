package main

import (
	"fmt"
	"io/ioutil"
)


/*  1、switch语句中case不需要写break
**	2、switch后可以没有表达式，直接在case里写判断条件
**/

func grade (score int) string {
	gd := ""
	switch {
	case score < 0 || score > 100:
		panic(fmt.Sprintf("Wrong score %d", score))
	case score < 60:
		gd = "F"
	case score < 80:
		gd = "C"
	case score < 90:
		gd = "B"
	case score <= 100:
		gd = "A"
	}
	return gd
}


/*  1、if条件里不需要括号
**	2、if条件里可以赋值
**	3、if条件里赋值的变量作用域就在这个if语句里
**/

func read() {
	const filename = "aaa.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n", contents)
	}
}


func main() {
	read()
	//fmt.Println(
	//	grade(0),
	//	grade(59),
	//	grade(60),
	//	grade(100),
	//	//grade(101),
	//)
}
