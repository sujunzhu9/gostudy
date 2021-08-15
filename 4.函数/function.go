package main

import "fmt"

//返回单个返回值
func foo1(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100
	return c
}

//返回多个返回值，匿名的
func foo2(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	return 111, 222
}

//返回多个返回值，有形参名称的
func foo3(a string, b int) (r1 int, r2 int) {
	fmt.Println("---foo3---")
	fmt.Println(a)
	fmt.Println(b)

	//给有名称的返回值变量赋值
	r1 = 1000
	r2 = 2000
	return
}

func foo4(a string, b int) (r1, r2 int) {
	fmt.Println("---foo4---")
	fmt.Println(a)
	fmt.Println(b)

	//给有名称的返回值变量赋值
	r1 = 10
	r2 = 20
	return
}

func main() {
	c := foo1("abc", 555)
	fmt.Println("c = ", c)

	ret1, ret2 := foo2("hhh", 222)
	fmt.Println(ret1, ret2)
	ret3, ret4 := foo3("aaa", 400)
	fmt.Println(ret3, ret4)
	ret5, ret6 := foo4("bbb", 200)
	fmt.Println(ret5, ret6)
}
