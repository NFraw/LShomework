package main

import (
	"fmt"
)

func main() {
	var a, b int       //定义变量
	println("请输入两个数字") //提示用户输入数字
	fmt.Scanf("%d %d", &a, &b)
	sum := a + b
	println("两数相加的结果是：", sum)
}
