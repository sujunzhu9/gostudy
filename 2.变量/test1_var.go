package main

import "fmt"

/*
	四种变量的声明方式
*/

//声明全局变量 方法一、二、三可用
//:=只能在函数体内使用

func main() {
	//方法一：声明一个变量
	var a int
	fmt.Println(a)

	//方法二：声明一个变量。初始化一个值
	var b int = 100
	fmt.Println(b)

	//方法三：初始化时候，可以省去数据类型，通过值来自动匹配
	var c = 100.01
	fmt.Println(c)
	fmt.Printf("Type:%T\n", c)

	//方法四：省去var直接自动匹配
	d := 100
	fmt.Println(d)
	fmt.Printf("Type:%T\n", d)

	//声明多个变量
	var xx, yy, aa, bb = 100, 200, "aa", "bb"
	fmt.Println(xx, yy, aa, bb)

	//多行声明
	var (
		l1 int     = 100
		l2 float64 = 12.09
	)
	fmt.Println(l1, l2)

}
