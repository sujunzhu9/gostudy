package main

import "fmt"

//const定义枚举类型
const (
	//iota每行加一
	one = iota
	two
	three
	four
)

const (
	a, b = iota + 1, iota + 2 //iota=0,a=iota+1,b=iota+2
	c, d                      //iota=1,c=iota+1,d=iota+2
	e, f                      //iota=2,e=iota+1,f=iota+2

	g, h = iota * 2, iota * 3 //iota=3,g=iota*2,h=iota*3
	i, j
)

func main() {
	//常量(只读属性)
	const length int = 10
	fmt.Println(length)
	fmt.Println(one, two, three, four)
	fmt.Println(a, b, c, d, e, f)
	fmt.Println(g, h, i, j)
}
