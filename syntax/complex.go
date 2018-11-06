package syntax

import (
	"fmt"
	"math"
	"math/cmplx"
)
// 支持复数类型
// float32, float64, complex64, complex128
// 验证欧拉公式
func euler () {
	// 已e为底，iπ为指数
	//fmt.Println(cmplx.Exp( 1i * math.Pi) + 1)
	//fmt.Println(cmplx.Pow(math.E, 1i * math.Pi) + 1)
	fmt.Printf("%.3f", cmplx.Exp( 1i * math.Pi) + 1)
}

func main() {
	euler()
}
