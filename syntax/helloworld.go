package main
import "fmt"

var v = 1 // 作用域是包package内部，不是全局的作用域
// var()集中定义变量
var (
	aa = 1
	bb = 2
	cc = "333"
)

// 只声明未赋值的变量
func variableZeroValue() {
	var a int
	var s string
	//fmt.Println(a, s)
	fmt.Printf("%d, %q \n", a, s)
}

// 初始化赋值变量
func variableInitialValue() {
	var a, b int = 3, 4
	var s string = "asdasdasd"
	println(a, b, s)
}

// 推断变量类型（无需声明变量类型）
func variableTypeDeduction() {
	var a, b, c, d = 3, 4, true, "asd"
	println(a, b, c, d)
}

// := 声明变量赋值，简化写法，不要重复定义变量，函数外面不能使用:=简化写法
func variableShorter() {
	a, b, c := 1, 3, 5
	fmt.Println(a, b, c)
}

func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	println(aa, bb, cc)
}
